package users

import (
	"github.com/gin-gonic/gin"

	"github.com/dmcalpin/go-cms/util/crud"
)

func AddRouter(r *gin.RouterGroup) {
	userCrud := crud.New[*User]()

	rg := r.Group("/api/users")
	rg.POST("/", userCrud.Create)
	rg.DELETE("/:key", userCrud.Delete)
	rg.PUT("/:key", userCrud.Update)
	rg.GET("/multi/:keys", userCrud.GetMulti)
	rg.GET("/:key", userCrud.Get)
	rg.GET("/", userCrud.GetAll)

	htmlRg := r.Group("/users")
	htmlRg.GET("/new", userCrud.CreateHTML)
	htmlRg.GET("/:key", userCrud.GetOneHTML)
	htmlRg.GET("/", userCrud.GetAllHTML)

}
