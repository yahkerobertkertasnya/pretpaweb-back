package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/yahkerobertkertasnya/preweb/database"
	"github.com/yahkerobertkertasnya/preweb/helper"
	"github.com/yahkerobertkertasnya/preweb/middleware"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
	"github.com/yahkerobertkertasnya/preweb/graph"
)

const defaultPort = "8080"

func main() {
	router := chi.NewRouter()
	//router.Use(middleware.Logger)

	cor := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost", "http://localhost:3000", "http://localhost:8080"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		Debug:            true,
	})

	router.Use(cor.Handler)
	//router.Get("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("Hello World!"))
	//})

	//http.ListenAndServe(":8080", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	database.MigrateTable()

	c := graph.Config{Resolvers: &graph.Resolver{
		DB: database.GetInstance(),
	}}

	c.Directives.Auth = func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		return helper.AuthDirectives(ctx, next)
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))

	router.Use(middleware.AuthMiddleware)
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
