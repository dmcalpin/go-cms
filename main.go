package main

import (
	"github.com/dmcalpin/go-cms/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := router.GetRouter()

	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	// r.TrustedPlatform = gin.PlatformGoogleAppEngine
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
