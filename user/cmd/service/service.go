package service

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"net"
	http2 "net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/prometheus"

	"user/pkg/endpoint"
	"user/pkg/grpc"
	"user/pkg/grpc/pb"
	http1 "user/pkg/http"
	"user/pkg/repo"
	"user/pkg/service"

	endpoint1 "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	lightsteptracergo "github.com/lightstep/lightstep-tracer-go"
	"github.com/oklog/oklog/pkg/group"
	opentracinggo "github.com/opentracing/opentracing-go"
	zipkingoopentracing "github.com/openzipkin-contrib/zipkin-go-opentracing"
	zipkingo "github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/reporter/http"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	grpc1 "google.golang.org/grpc"
	"sourcegraph.com/sourcegraph/appdash"
	"sourcegraph.com/sourcegraph/appdash/opentracing"
)

var tracer opentracinggo.Tracer
var logger log.Logger

var fs = flag.NewFlagSet("user", flag.ExitOnError)
var debugAddr = fs.String("debug-addr", ":8080", "Debug and metrics listen address")
var httpAddr = fs.String("http-addr", ":8081", "HTTP listen address")
var grpcAddr = fs.String("grpc-addr", ":8082", "gRPC listen address")
var zipkinURL = fs.String("zipkin-url", "", "Enable Zipkin tracing via a collector URL e.g. http://localhost:9411/api/v1/spans")
var lightstepToken = fs.String("lightstep-token", "", "Enable LightStep tracing via a LightStep access token")
var appdashAddr = fs.String("appdash-addr", "", "Enable Appdash tracing via an Appdash server host:port")

func Run() {
	_ = fs.Parse(os.Args[1:])

	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	if *zipkinURL != "" {
		_ = logger.Log("tracer", "Zipkin", "URL", *zipkinURL)
		reporter := http.NewReporter(*zipkinURL)
		defer reporter.Close()
		enp, err := zipkingo.NewEndpoint("user", "localhost:80")
		if err != nil {
			_ = logger.Log("err", err)
			os.Exit(1)
		}
		localEndpoint := zipkingo.WithLocalEndpoint(enp)
		nativeTracer, err := zipkingo.NewTracer(reporter, localEndpoint)
		if err != nil {
			_ = logger.Log("err", err)
			os.Exit(1)
		}
		tracer = zipkingoopentracing.Wrap(nativeTracer)
	} else if *lightstepToken != "" {
		_ = logger.Log("tracer", "LightStep")
		tracer = lightsteptracergo.NewTracer(lightsteptracergo.Options{AccessToken: *lightstepToken})
		defer lightsteptracergo.Flush(context.Background(), tracer)
	} else if *appdashAddr != "" {
		_ = logger.Log("tracer", "Appdash", "addr", *appdashAddr)
		collector := appdash.NewRemoteCollector(*appdashAddr)
		tracer = opentracing.NewTracer(collector)
		defer collector.Close()
	} else {
		_ = logger.Log("tracer", "none")
		tracer = opentracinggo.GlobalTracer()
	}

	db, err := sql.Open("postgres", "postgres://postgres:secret@localhost:10521/user_service?sslmode=disable")
	if err != nil {
		panic("invalid database connection source")
	}
	if db.Ping() != nil {
		panic("failed to ping database")
	}
	queries := repo.New(db)

	svc := service.New(queries, getServiceMiddleware(logger))
	eps := endpoint.New(svc, getEndpointMiddleware(logger))
	g := createService(eps)
	initMetricsEndpoint(g)
	initCancelInterrupt(g)
	_ = logger.Log("exit", g.Run())

}

func initHttpHandler(endpoints endpoint.Endpoints, g *group.Group) {
	options := defaultHttpOptions(logger, tracer)

	httpHandler := http1.NewHTTPHandler(endpoints, options)
	httpListener, err := net.Listen("tcp", *httpAddr)
	if err != nil {
		_ = logger.Log("transport", "HTTP", "during", "Listen", "err", err)
	}
	g.Add(func() error {
		_ = logger.Log("transport", "HTTP", "addr", *httpAddr)
		return http2.Serve(httpListener, httpHandler)
	}, func(error) {
		_ = httpListener.Close()
	})

}

func getLatency() metrics.Histogram {
	return prometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "myteam",
		Subsystem: "foosvc",
		Name:      "request_latency_seconds",
		Help:      "Incoming request latency in seconds.",
	}, []string{"method", "status_code"})
}

func getServiceMiddleware(_ log.Logger) (mw []service.Middleware) {
	return
}

func getEndpointMiddleware(logger log.Logger) (mw map[string][]endpoint1.Middleware) {
	mw = map[string][]endpoint1.Middleware{}
	addEndpointMiddlewareToAllMethods(getEndpointMiddleware(logger), endpoint.LoggingMiddleware(logger))
	addEndpointMiddlewareToAllMethods(getEndpointMiddleware(logger), endpoint.InstrumentingMiddleware(getLatency()))

	return
}

func initMetricsEndpoint(g *group.Group) {
	http2.DefaultServeMux.Handle("/metrics", promhttp.Handler())
	debugListener, err := net.Listen("tcp", *debugAddr)
	if err != nil {
		_ = logger.Log("transport", "debug/HTTP", "during", "Listen", "err", err)
	}
	g.Add(func() error {
		_ = logger.Log("transport", "debug/HTTP", "addr", *debugAddr)
		return http2.Serve(debugListener, http2.DefaultServeMux)
	}, func(error) {
		_ = debugListener.Close()
	})
}

func initCancelInterrupt(g *group.Group) {
	cancelInterrupt := make(chan struct{})
	g.Add(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-c:
			return fmt.Errorf("received signal %s", sig)
		case <-cancelInterrupt:
			return nil
		}
	}, func(error) {
		close(cancelInterrupt)
	})
}

func initGRPCHandler(endpoints endpoint.Endpoints, g *group.Group) {
	options := defaultGRPCOptions(logger, tracer)

	grpcServer := grpc.NewGRPCServer(endpoints, options)
	grpcListener, err := net.Listen("tcp", *grpcAddr)
	if err != nil {
		_ = logger.Log("transport", "gRPC", "during", "Listen", "err", err)
	}
	g.Add(func() error {
		_ = logger.Log("transport", "gRPC", "addr", *grpcAddr)
		baseServer := grpc1.NewServer()
		pb.RegisterUserServer(baseServer, grpcServer)
		return baseServer.Serve(grpcListener)
	}, func(error) {
		_ = grpcListener.Close()
	})

}
