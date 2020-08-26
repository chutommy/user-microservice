package server

import (
	"net/http"

	"github.com/chutified/appointments/accounts/config"
	"github.com/chutified/appointments/accounts/controllers"
	"github.com/pkg/errors"
)

// Server allows server's routings and the initialization.
type Server interface {
	Init(*config.Config) error
}

type server struct{}

// New constructs the Server interface value.
func New() Server {
	return &server{}
}

func (s *server) Init(cfg *config.Config) error {

	// initialize handler
	h := controllers.New()
	err := h.Init(cfg.Db)
	if err != nil {
		return errors.Wrap(err, "initializing handler")
	}
	defer h.Close()

	// set routings
	r := SetRoutes(h)

	// setup the server
	srv := http.Server{
		Addr:              cfg.Server.Address,
		Handler:           r,
		ReadTimeout:       cfg.Server.ReadTimeout,
		ReadHeaderTimeout: cfg.Server.ReadHeaderTimeout,
		WriteTimeout:      cfg.Server.WriteTimeout,
		IdleTimeout:       cfg.Server.IdleTimeout,
		MaxHeaderBytes:    http.DefaultMaxHeaderBytes,
	}

	// run
	// TODO gracefully shutdown
	err = srv.ListenAndServe()
	if err != nil {
		return errors.Wrap(err, "running the server")
	}
	return nil
}
