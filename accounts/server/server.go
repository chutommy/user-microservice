package server

import (
	"net/http"

	"github.com/chutified/booking-terminal/accounts/config"
	"github.com/chutified/booking-terminal/accounts/controllers"
	"github.com/pkg/errors"
)

// Server allows server's routings and the initialization.
type Server interface {
	Init(*config.Config) error
}

type server struct {
	h   *controllers.Handler
	srv *http.Server
}

// New constructs the Server interface value.
func New() Server {
	return &server{}
}

// Init initializes server.
func (s *server) Init(cfg *config.Config) error {

	// initialize handler
	s.h = controllers.New()
	err := s.h.Init(cfg.Db)
	if err != nil {
		return errors.Wrap(err, "initializing handler")
	}

	// set routings
	r := SetRoutes(s.h)

	// setup the server
	s.srv = &http.Server{
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
