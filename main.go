package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	db "github.com/brkss/golang-graphql/db/sqlc"
	"github.com/brkss/golang-graphql/directives"
	"github.com/brkss/golang-graphql/graph"
	"github.com/brkss/golang-graphql/utils"
	_ "github.com/lib/pq"
)

const defaultPort = "8080"

const (
	dbDriver = "postgres"
	dbSource = "postgres://root:root@localhost:5432/gqlvf?sslmode=disable"
)

func main() {
	

	// load config 
	config, err := utils.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config : ", err)
	}

	fmt.Printf("%+v\n", config)

	con, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to database : ", err)
	}

	store := db.NewStore(con)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	c := graph.Config{Resolvers: &graph.Resolver{ Store: store }}
	c.Directives.Binding = directives.Binding

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
