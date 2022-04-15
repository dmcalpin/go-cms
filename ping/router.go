package ping

import (
	"github.com/gin-gonic/gin"
)

func AddRouter(r *gin.RouterGroup) {
	rg := r.Group("/ping")

	rg.GET("/", getPing)
	rg.GET("/html", getHtml)
}
