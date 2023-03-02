package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/koralle/lazy-warehouse-backend-prototype/config"
	"github.com/koralle/lazy-warehouse-backend-prototype/mux"
)

func run() error {
	config, err := config.New()
	if err != nil {
		return err
	}
	mux, cleanup, err := mux.NewMux()

	if err != nil {
		return err
	}

	defer cleanup()

	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), mux)

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Printf("failed to terminate server: %v", err)
		os.Exit(1)
	}
}
