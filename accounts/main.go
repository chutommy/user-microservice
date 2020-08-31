package main

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/chutified/booking-terminal/accounts/config"
	"github.com/chutified/booking-terminal/accounts/server"
	_ "github.com/lib/pq"
)

func main() {

	// get config
	cfg, err := config.GetConfig("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	// handle gracefull shutdown
	errs := make(chan error)
	go func() {

		sig := make(chan os.Signal)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

		errs <- errors.New((<-sig).String())
	}()

	// init the server
	srv := server.New()
	go func() {

		log.Printf("Server launching on %s...\n", cfg.Server.Address)
		err = srv.Init(cfg)
		if err != nil {
			log.Fatal(err)
		}

		err = srv.Start()
		if err != nil {
			log.Fatal(err)
		}
	}()
	defer func() {

		err = srv.Stop()
		if err != nil {
			log.Println(err)
		}

		log.Println("Server succesfully shut down.")
	}()

	log.Printf("Server gracefully shutting down... (%s)\n", <-errs)
}
