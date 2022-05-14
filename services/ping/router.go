package ping

import (
	"github.com/gin-gonic/gin"
)

func AddRouter(r *gin.RouterGroup) {
	rg := r.Group("/api/ping")
	rg.GET("/", getPing)

	htmlRg := r.Group("/ping")
	htmlRg.GET("/", getHtml)
}
