package router

import (
	"html/template"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/dmcalpin/go-cms/services/ping"
	"github.com/dmcalpin/go-cms/services/users"
)

var router *gin.Engine

func init() {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	tmpl, err := template.ParseFiles("services/ping/ping.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	r.SetHTMLTemplate(tmpl)

	rg := r.Group("/api")
	ping.AddRouter(rg)
	users.AddRouter(rg)

	router = r
}

func GetRouter() *gin.Engine {
	return router
}
