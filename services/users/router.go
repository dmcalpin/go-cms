package users

import (
	"github.com/gin-gonic/gin"

	"cloud.google.com/go/datastore"

	"github.com/dmcalpin/go-cms/db"

	crud "github.com/dmcalpin/go-cms/util/crud"
)

const UserKind = "User"

type User struct {
	Key          *datastore.Key `datastore:"__key__" json:"key"`
	EmailAddress string         `datastore:"emailAddress" json:"emailAddress"`
	Password     string         `datastore:"password" json:"password"`
}

type UserCreateInput struct {
	EmailAddress string `datastore:"emailAddress" json:"emailAddress" binding:"required,email"`
	Password     string `datastore:"password" json:"password" binding:"required"`
}

type UserUpdateInput struct {
	EmailAddress string `datastore:"emailAddress,omitempty" json:"emailAddress,omitempty" binding:"email"`
	Password     string `datastore:"password,omitempty" json:"password,omitempty"`
}

func AddRouter(r *gin.RouterGroup) {
	rg := r.Group("/users")

	userCrud := crud.New[User, UserCreateInput, UserUpdateInput](UserKind, db.GetClient())

	rg.POST("/", userCrud.Create)
	rg.DELETE("/:key", userCrud.Delete)
	rg.PUT("/:key", userCrud.Update)
	rg.GET("/:key", userCrud.Get)
	rg.GET("/", userCrud.GetAll)

}
