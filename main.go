package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/aakazanskayaa/comments-post-system/config"
	"github.com/aakazanskayaa/comments-post-system/db"
	"github.com/aakazanskayaa/comments-post-system/internal/graph"
)

func main() {
	// Загружаем конфиг
	cfg := config.LoadConfig()

	// Инициализируем хранилище
	db.InitStorage(cfg)

	// GraphQL-сервер
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	// маршруты
	http.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	http.Handle("/query", srv)

	// порт (по умолчанию 8080)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on http://localhost:%s/\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
