package crud

import (
	"net/http"

	"github.com/dmcalpin/go-cms/db"
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
		c.JSON(crud.errToHTTPError(err), nil)
		return
	}

	key, err := db.Client.Put(c, datastore.IncompleteKey(crud.Kind, nil), input)
	if err != nil {
		c.Error(err)
		c.JSON(crud.errToHTTPError(err), nil)
		return
	}

	var entity T
	// .New method isn't great, but it's not possible
	// to do T{}, and the .Get below needs a pointer
	// to a zero value struct
	e := entity.New()
	err = db.Client.Get(c, key, e)

	c.JSON(http.StatusCreated, e)
}
