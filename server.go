package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Roshantwanabasu/gql-in-go-demo/graph"
	"github.com/Roshantwanabasu/gql-in-go-demo/graph/generated"
	"github.com/Roshantwanabasu/gql-in-go-demo/internal/auth"
	database "github.com/Roshantwanabasu/gql-in-go-demo/internal/pkg/db/migrations/mysql"
	"github.com/go-chi/chi/v5"
)

const defaultPort = "8000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	router := chi.NewRouter()
	router.Use(auth.Middleware())

	database.InitDB()
	database.Migrate()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
