package users

import (
	"github.com/gin-gonic/gin"
)

func AddRouter(r *gin.RouterGroup) {
	rg := r.Group("/users")

	rg.POST("/", createUser)
	rg.DELETE("/:key", deleteUser)
	rg.PUT("/:key", updateUser)
	rg.GET("/:key", getUser)
	rg.GET("/", getAllUsers)

}
