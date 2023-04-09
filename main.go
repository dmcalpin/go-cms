package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/dmcalpin/go-cms/services/jobs"
	"github.com/dmcalpin/go-cms/services/ping"
	"github.com/dmcalpin/go-cms/services/users"
)

func main() {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile("./static", false)))

	rg := r.Group("/")
	ping.AddRouter(rg)
	users.AddRouter(rg)
	jobs.AddRouter(rg)

	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	// r.TrustedPlatform = gin.PlatformGoogleAppEngine
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
