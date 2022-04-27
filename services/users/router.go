package users

import (
	"context"
	"encoding/json"

	"cloud.google.com/go/datastore"
	"github.com/gin-gonic/gin"

	"github.com/dmcalpin/go-cms/db"
	"github.com/dmcalpin/go-cms/services/jobs"
	"github.com/dmcalpin/go-cms/util/crud"
)

const UserKind = "User"

type User struct {
	Key          *datastore.Key `datastore:"__key__" json:"key"`
	EmailAddress string         `datastore:"emailAddress" json:"emailAddress"`
	Password     string         `datastore:"password" json:"password"`
	Job          *datastore.Key `datastore:"job" json:"job"`
}

func (u *User) MarshalJSON() ([]byte, error) {

	userData := map[string]interface{}{
		"key":          u.Key.Encode(),
		"emailAddress": u.EmailAddress,
		"job":          u.Job,
	}

	if u.Job != nil {
		job := jobs.Job{}
		client := db.GetClient()
		err := client.Get(context.Background(), u.Job, &job)
		if err != nil {
			return nil, err
		}

		userData["job"] = map[string]interface{}{
			"title":      job.Title,
			"descriptin": job.Description,
		}
	}

	return json.Marshal(userData)
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

	userCrud := crud.New[User, UserCreateInput, UserUpdateInput](UserKind, db.GetClient())

	rg.POST("/", userCrud.Create)
	rg.DELETE("/:key", userCrud.Delete)
	rg.PUT("/:key", userCrud.Update)
	rg.GET("/:key", userCrud.Get)
	rg.GET("/", userCrud.GetAll)

}
