package auth

import (
	"github.com/gin-gonic/gin"
)

func AddRouter(r *gin.RouterGroup) *gin.RouterGroup {
	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	authorized := r.Group("/api/auth")
	authorized.Use(gin.BasicAuth(gin.Accounts{
		"foo":  "bar",
		"manu": "123",
	}))
	rg := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	return rg
}
