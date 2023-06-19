package initializer

import (
	"encoding/gob"
	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	. "gogo-boot/go-boot/platform/config"
	"gogo-boot/go-boot/platform/middleware"
	"golang.org/x/oauth2"
)

var Router *gin.Engine

func init() {

	logLevel, _ := log.ParseLevel(AppConfig.Server.LogLevel)
	log.SetLevel(logLevel)
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{}) //for easy parsing by logstash or Splunk

	// HTTP Server Set up
	// Router = gin.Default() // Default Mode
	Router = gin.New()
	Router.Use(gin.Recovery())

	// Create Casbin Enforcer
	e, _ := casbin.NewEnforcer("/config/authz_model.conf", "/config/authz_policy.csv")

	Router.Use(middleware.LoggingMiddleware(), middleware.NewAuthorizer(e))

	// Load HTML Template
	Router.LoadHTMLGlob("web/template/*")
	Router.Static("/public", "web/static")

	var profile struct {
		FamilyName string   `json:"family_name"`
		GivenName  string   `json:"given_name"`
		Groups     []string `json:"groups"`
		Email      string   `json:"email"`
		Name       string   `json:"name"`
		Roles      []string `json:"roles"`
	}

	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(profile)
	gob.Register(oauth2.Token{})

	store := memstore.NewStore([]byte("secret"))
	Router.Use(sessions.Sessions("auth-session", store))
}
