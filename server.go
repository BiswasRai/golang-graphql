package main

import (
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/biswasRai/golang-graphql/graph"
	"github.com/biswasRai/golang-graphql/internal/database"
	"github.com/biswasRai/golang-graphql/internal/middleware"
	"github.com/gin-gonic/gin"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db := database.ConnectDb()

	server := gin.Default()

	server.Use(middleware.Auth())

	server.POST("/query", func(ctx *gin.Context) {
		handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{DB: db}})).ServeHTTP(ctx.Writer, ctx.Request)
	})

	server.GET("/", func(ctx *gin.Context) {
		playground.Handler("GraphQL", "/query").ServeHTTP(ctx.Writer, ctx.Request)

	})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	server.Run()

}
