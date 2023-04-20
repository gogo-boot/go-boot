package main

import (
	. "example.com/go-web-template/config"
	"example.com/go-web-template/graph"
	"example.com/go-web-template/middlewares"
	myOidc "example.com/go-web-template/oidc"
	"example.com/go-web-template/openapi"
	"example.com/go-web-template/restapi"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var server *gin.Engine

func init() {
	logLevel, _ := log.ParseLevel(AppConfig.Server.LogLevel)
	log.SetLevel(logLevel)
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{}) //for easy parsing by logstash or Splunk

	// HTTP Server Set up
	// server = gin.Default() // Default Mode
	server = gin.New()
	server.Use(gin.Recovery())
	server.Use(middlewares.LoggingMiddleware())
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {

	server.POST("/query", graphqlHandler())
	server.GET("/", playgroundHandler())
	openapi.NewRouter(server.Group("/openapi"))
	restapi.Routes(server.Group("/restapi"))
	//graphql.Routes(server.Group("/graphql"))
	//oauth2.Routes(server.Group("/login"))
	myOidc.Routes(server.Group("/login"))

	server.Run(":" + AppConfig.Server.PortNumber)
}
