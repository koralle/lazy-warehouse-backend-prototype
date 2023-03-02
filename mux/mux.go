package mux

import (
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/koralle/lazy-warehouse-backend-prototype/graph"
)

func NewMux() (http.Handler, func(), error) {
	mux := chi.NewRouter()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	mux.Post("/query", srv.ServeHTTP)
	mux.Get("/", playground.Handler("GraphQL playground", "/query"))

	return mux, func() { fmt.Println("Clean up.") }, nil

}
