package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/koralle/lazy-warehouse-backend-prototype/config"
	"github.com/koralle/lazy-warehouse-backend-prototype/mux"
	"github.com/koralle/lazy-warehouse-backend-prototype/server"
)

func run(ctx context.Context) error {
	config, err := config.New()
	if err != nil {
		return err
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		log.Fatalf("failed to listen port %d: %v", config.Port, err)
	}

	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("start with: %v", url)

	mux, cleanup, err := mux.NewMux()

	if err != nil {
		return err
	}

	defer cleanup()

	s := server.NewServer(l, mux)

	return s.Run(ctx)

}

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminate server: %v", err)
		os.Exit(1)
	}
}
