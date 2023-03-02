package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/koralle/lazy-warehouse-backend-prototype/config"
	"github.com/koralle/lazy-warehouse-backend-prototype/graph"
)

const defaultPort = "8080"

func run() error {
	cfg, err := config.New()
	if err != nil {
		return err
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%d/ for GraphQL playground", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+fmt.Sprintf("%d", cfg.Port), nil))

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Printf("failed to terminate server: %v", err)
		os.Exit(1)
	}
}
