package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	db "github.com/brkss/golang-graphql/db/sqlc"
	"github.com/brkss/golang-graphql/graph"
	_ "github.com/lib/pq"
)

const defaultPort = "8080"

const (
	dbDriver = "postgres"
	dbSource = "postgres://root:root@localhost:5432/gqlvf?sslmode=disable"
)

func main() {
	

	con, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to database : ", err)
	}

	store := db.NewStore(con)


	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		Store: store,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
