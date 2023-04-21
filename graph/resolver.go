package graph

//go:generate go run github.com/99designs/gqlgen generate

import (
	"example.com/go-boot/graph/model"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
type Resolver struct {
	todos []*model.Todo
}

func Routes(rg *gin.RouterGroup) {
	rg.POST("/query", graphqlHandler())
	rg.GET("/", playgroundHandler())
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: &Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
