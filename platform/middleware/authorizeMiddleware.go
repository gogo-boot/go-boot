package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

// NewAuthorizer returns the authorizer, uses a Casbin enforcer as input
func NewAuthorizer(e *casbin.Enforcer) gin.HandlerFunc {
	a := &BasicAuthorizer{enforcer: e}

	return func(c *gin.Context) {
		if !a.CheckPermission(c.Request) {
			a.RequirePermission(c)
		}
	}
}

// BasicAuthorizer stores the casbin handler
type BasicAuthorizer struct {
	enforcer *casbin.Enforcer
}

// GetUserName gets the user name from the request.
// Currently, only HTTP basic authentication is supported
func (a *BasicAuthorizer) GetUserName(r *http.Request) string {
	username, _, _ := r.BasicAuth()
	return username
}

// CheckPermission checks the user/method/path combination from the request.
// Returns true (permission granted) or false (permission forbidden)
func (a *BasicAuthorizer) CheckPermission(r *http.Request) bool {
	user := a.GetUserName(r)
	method := r.Method
	path := r.URL.Path

	allowed, err := a.enforcer.Enforce(user, path, method)
	if err != nil {
		panic(err)
	}

	return allowed
}

// RequirePermission returns the 403 Forbidden to the client
func (a *BasicAuthorizer) RequirePermission(c *gin.Context) {
	c.AbortWithStatus(http.StatusForbidden)
}

//func xxx() {
//
//	e, err := casbin.NewEnforcer("authz_model.conf", "authz_policy.csv")
//
//	sub := "alice" // the user that wants to access a resource.
//	obj := "data1" // the resource that is going to be accessed.
//	act := "read"  // the operation that the user performs on the resource.
//
//	ok, err := e.Enforce(sub, obj, act)
//
//	if err != nil {
//		// handle err
//	}
//
//	if ok == true {
//		// permit alice to read data1
//	} else {
//		// deny the request, show an error
//	}
//
//	// You could use BatchEnforce() to enforce some requests in batches.
//	// This method returns a bool slice, and this slice's index corresponds to the row index of the two-dimensional array.
//	// e.g. results[0] is the result of {"alice", "data1", "read"}
//	results, err := e.BatchEnforce([][]interface{}{{"alice", "data1", "read"}, {"bob", "data2", "write"}, {"jack", "data3", "read"}})
//
//	roles, err := e.GetRolesForUser("alice")
//}
