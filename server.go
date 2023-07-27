package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	//"github.com/graphql-go/handler"
	"github.com/rs/cors"
	"github.com/yahkerobertkertasnya/preweb/database"
	"github.com/yahkerobertkertasnya/preweb/graph"
	"github.com/yahkerobertkertasnya/preweb/helper"
	"github.com/yahkerobertkertasnya/preweb/middleware"
	"log"
	"net/http"
	"os"
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

	srv := handler.New(graph.NewExecutableSchema(c))
	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})
	srv.SetQueryCache(lru.New(1000))
	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	router.Use(middleware.AuthMiddleware)
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
