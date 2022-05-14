package users

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"cloud.google.com/go/datastore"
	"github.com/gin-gonic/gin"

	"github.com/dmcalpin/go-cms/db"
	"github.com/dmcalpin/go-cms/util/crud"
)

const UserKind = "User"

type User struct {
	db.DatastoreModel
	EmailAddress string         `datastore:"emailAddress" json:"emailAddress"`
	Password     string         `datastore:"password" json:"-"`
	Job          *datastore.Key `datastore:"job" json:"job"`
}

func (u *User) Patch(i interface{}) {
	input := i.(*User)
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

func (u *User) Get(c context.Context) error {
	return db.Client.Get(c, u.Key, u)
}

func (u *User) Save(c context.Context) error {
	updatedKey, err := db.Client.Put(c, u.Key, u)
	if err != nil {
		return err
	}

	u.Key = updatedKey

	return nil
}

func (u *User) SaveAndGet(c context.Context) error {
	err := u.Save(c)
	if err != nil {
		return err
	}

	return u.Get(c)
}

func (u *User) Delete(c context.Context) error {
	return db.Client.Delete(c, u.Key)
}

func (u *User) New(key *datastore.Key) db.Patchable {
	user := &User{}
	user.Kind = UserKind
	user.Key = key

	return user
}

func (u *User) Validate() error {
	if !strings.Contains(u.EmailAddress, "@") {
		return errors.New("bad email address, must contain '@'")
	}
	return nil
}

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
	htmlRg.GET("/", func(c *gin.Context) {
		var users []*User
		query := datastore.NewQuery(UserKind)
		_, err := db.Client.GetAll(c, query, &users)
		if err != nil {
			c.Error(err)
			c.HTML(http.StatusBadRequest, "404.gohtml", map[string]interface{}{
				"Status":  http.StatusBadRequest,
				"Message": err.Error(),
			})
		}

		c.HTML(http.StatusOK, "users_list.gohtml", users)
	})
	htmlRg.GET("/:key", func(c *gin.Context) {
		keyStr, ok := c.Params.Get("key")
		if !ok {
			c.Error(errors.New("key param required"))
			c.HTML(http.StatusBadRequest, "404.gohtml", map[string]interface{}{
				"Status":  http.StatusBadRequest,
				"Message": "key is a required param",
			})
		}

		key, err := datastore.DecodeKey(keyStr)
		if err != nil {
			c.Error(errors.New("key param required"))
			c.HTML(http.StatusBadRequest, "404.gohtml", map[string]interface{}{
				"Status":  http.StatusBadRequest,
				"Message": err.Error(),
			})
		}

		user := &User{}
		err = db.Client.Get(c, key, &user)
		if err != nil {
			c.Error(err)
			c.HTML(http.StatusBadRequest, "404.gohtml", map[string]interface{}{
				"Status":  http.StatusBadRequest,
				"Message": err.Error(),
			})
		}

		c.HTML(http.StatusOK, "users_list.gohtml", user)
	})
}
