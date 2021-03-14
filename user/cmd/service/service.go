package service

import (
	"database/sql"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	gmdw "github.com/grpc-ecosystem/go-grpc-middleware"
	gzap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	gtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/chutified/booking-terminal/user/pkg/grpc/userpb"
	"github.com/chutified/booking-terminal/user/pkg/repo"
	"github.com/chutified/booking-terminal/user/pkg/service"
)

var fs = flag.NewFlagSet("user", flag.ExitOnError)
var debugMode = fs.Bool("debug", false, "enable development level logging")
var _ = fs.String("http-port", "8081", "HTTP listen address port")
var grpcPort = fs.String("grpc-port", "8082", "gRPC listen address port")
var dbURL = fs.String("db_url", "", "database URL of the user service")

func Run() (err error) {
	err = fs.Parse(os.Args[1:])
	if err != nil {
		return err
	}

	// create a structured logger
	var logger *zap.Logger
	if *debugMode {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		return err
	}
	defer func() {
		err = logger.Sync()
	}()

	// connect to the DB
	dbLog := logger.With(zap.String("db_conn_url", *dbURL))
	var attempts int8 = 3
	var db *sql.DB
	for a := int8(1); a <= attempts; a++ {
		time.Sleep(3000 * time.Millisecond)

		db, err = sql.Open("postgres", *dbURL)
		if err != nil {
			dbLog.Warn(
				"failed to open database",
				zap.Int8("attempt", a),
				zap.Error(err),
			)
		} else {
			break
		}
	}
	if err != nil {
		dbLog.Fatal(
			"failed to connect to the database after the max number of attempts",
			zap.Int8("max_attempts", attempts),
			zap.Error(err),
		)
	}
	// test the DB connection
	if db.Ping() != nil {
		dbLog.Fatal(
			"failed to ping the database",
			zap.Error(err),
		)
	}
	dbLog.Info("successfully connected to the database")

	// construct a repo
	qrs := repo.New(db)

	// build a logger interceptor middleware
	opts := []gzap.Option{}

	// init user service's server
	userSrv := service.NewUserServer(qrs)
	grpcSrv := grpc.NewServer(
		gmdw.WithUnaryServerChain(
			gtags.UnaryServerInterceptor(gtags.WithFieldExtractor(gtags.CodeGenRequestFieldExtractor)),
			gzap.UnaryServerInterceptor(logger, opts...),
		),
	)
	userpb.RegisterUserServiceServer(grpcSrv, userSrv)
	reflection.Register(grpcSrv)

	// listen
	address := fmt.Sprintf("0.0.0.0:%s", *grpcPort)
	l, err := net.Listen("tcp", address)
	if err != nil {
		logger.Fatal(
			"failed to listen the network",
			zap.String("address", address),
			zap.Error(err),
		)
	}

	// server
	go func() {
		logger.Info("user service's server online")

		if err = grpcSrv.Serve(l); err != nil {
			logger.Fatal(
				"server shutdown err",
				zap.Error(err),
			)
		}
	}()

	// listen for signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	sig := <-c
	logger.Info("received terminate signal", zap.String("signal", sig.String()))
	close(c)

	return nil
}
