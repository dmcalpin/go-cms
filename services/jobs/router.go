package jobs

import (
	"github.com/gin-gonic/gin"

	"github.com/dmcalpin/go-cms/util/crud"
)

func AddRouter(r *gin.RouterGroup) {
	jobCrud := crud.New[*Job]()

	rg := r.Group("/api/jobs")
	rg.POST("/", jobCrud.Create)
	rg.DELETE("/:key", jobCrud.Delete)
	rg.PUT("/:key", jobCrud.Update)
	rg.GET("/multi/:keys", jobCrud.GetMulti)
	rg.GET("/:key", jobCrud.Get)
	rg.GET("/", jobCrud.GetAll)

	htmlRg := r.Group("/jobs")
	htmlRg.GET("/:key", jobCrud.GetOneHTML)
	htmlRg.GET("/", jobCrud.GetAllHTML)
}
