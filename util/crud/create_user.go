package users

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"cloud.google.com/go/datastore"
)

func (crud *CRUD[T, T2, T3]) Create(c *gin.Context) {
	// _ = c.MustGet(gin.AuthUserKey).(string)

	// Parse JSON
	input := new(T2)
	err := c.Bind(input)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	key, err := crud.DB.Put(c, datastore.IncompleteKey(crud.Kind, nil), input)
	if err != nil {
		c.Error(err)
		c.JSON(crud.errToHTTPError(err), nil)
		return
	}

	entity := new(T)
	err = crud.DB.Get(c, key, entity)

	c.JSON(http.StatusCreated, &entity)
}
