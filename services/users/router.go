package users

import (
	"cloud.google.com/go/datastore"
	"github.com/gin-gonic/gin"

	"github.com/dmcalpin/go-cms/util/crud"
)

const UserKind = "User"

type User struct {
	Key          *datastore.Key `datastore:"__key__" json:"key"`
	EmailAddress string         `datastore:"emailAddress" json:"emailAddress"`
	Password     string         `datastore:"password" json:"-"`
	Job          *datastore.Key `datastore:"job" json:"job"`
}

func (u *User) Patch(i interface{}) {
	input := i.(*UserUpdateInput)
	if input.EmailAddress != "" {
		u.EmailAddress = input.EmailAddress
	}
	if input.Password != "" {
		u.Password = input.Password
	}
	if input.Job != nil {
		u.Job = input.Job
	}
}

func (u *User) New() crud.Patchable {
	return &User{}
}

type UserCreateInput struct {
	EmailAddress string         `datastore:"emailAddress" json:"emailAddress" binding:"required,email"`
	Password     string         `datastore:"password" json:"password" binding:"required"`
	Job          *datastore.Key `datastore:"job" json:"job"`
}

type UserUpdateInput struct {
	EmailAddress string         `datastore:"emailAddress,omitempty" json:"emailAddress,omitempty" binding:"email"`
	Password     string         `datastore:"password,omitempty" json:"password,omitempty"`
	Job          *datastore.Key `datastore:"job" json:"job"`
}

func AddRouter(r *gin.RouterGroup) {
	rg := r.Group("/users")

	userCrud := crud.New[*User, UserCreateInput, UserUpdateInput](UserKind)

	rg.POST("/", userCrud.Create)
	rg.DELETE("/:key", userCrud.Delete)
	rg.PUT("/:key", userCrud.Update)
	rg.GET("/multi/:keys", userCrud.GetMulti)
	rg.GET("/:key", userCrud.Get)
	rg.GET("/", userCrud.GetAll)
}
