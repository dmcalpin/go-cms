package main

import (
	"github.com/gin-gonic/gin"

	"github.com/dmcalpin/go-cms/services/jobs"
	"github.com/dmcalpin/go-cms/services/ping"
	"github.com/dmcalpin/go-cms/services/users"
)

func main() {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	ping.SetTemplates(r)

	rg := r.Group("/api")
	ping.AddRouter(rg)
	users.AddRouter(rg)
	jobs.AddRouter(rg)

	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	// r.TrustedPlatform = gin.PlatformGoogleAppEngine
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
