package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/aakazanskayaa/comments-post-system/internal/graph"
)

func main() {
	// Создаём GraphQL сервер
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{},
	}))

	// Настраиваем маршруты
	http.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	http.Handle("/query", srv)

	// Запускаем сервер
	fmt.Println("🚀 Server running on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
