package main

import (
	"fmt"
	"log"
	"net/http"
	"wiiki_server/infra/graph"
	"wiiki_server/infra/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {

	// TODO config
	// conf :=

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", 8080), nil))

}
