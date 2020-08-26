package controllers

import (
	"github.com/chutified/booking-terminal/accounts/config"
	"github.com/chutified/booking-terminal/accounts/data"
	"github.com/pkg/errors"
)

// Handler is the controller of the accounts.
type Handler struct {
	ds *data.DatabaseService
}

// New is the contructor for the controller.
func New() *Handler {
	h := &Handler{
		ds: data.New(),
	}

	return h
}

// Init initializes the Handler contructor.
func (h *Handler) Init(cfg *config.DBConfig) error {

	// initialize database service
	err := h.ds.Init(cfg)
	if err != nil {
		return errors.Wrap(err, "initializing database service")
	}

	return nil
}

// Close closes all connections of Handler's.
func (h *Handler) Close() error {

	// close database service
	err := h.ds.Stop()
	if err != nil {
		errors.Wrap(err, "stoping database service")
	}

	return nil
}
