package users

import (
	"errors"
	"net/http"

	"cloud.google.com/go/datastore"
	"github.com/gin-gonic/gin"

	"github.com/dmcalpin/go-cms/db"
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
