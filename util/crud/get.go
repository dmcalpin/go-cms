package crud

import (
	"fmt"
	"net/http"
	"strings"

	"cloud.google.com/go/datastore"
	"github.com/dmcalpin/go-cms/db"
	"github.com/gin-gonic/gin"
)

func (crud *CRUD[T, T2, T3]) Get(c *gin.Context) {
	key, err := crud.decodeKeyParam(&c.Params)
	if err != nil {
		crud.logAndWriteError(c, ErrInvalidInput)
		return
	}

	var entity T
	// .New method isn't great, but it's not possible
	// to do T{}, and the .Get below needs a pointer
	// to a zero value struct
	e := entity.New()
	err = db.Client.Get(c, key, e)
	if err != nil {
		crud.logAndWriteError(c, err)
		return
	}

	c.JSON(http.StatusOK, e)
}

func (crud *CRUD[T, T2, T3]) GetAll(c *gin.Context) {
	var entities []T

	query := datastore.NewQuery(crud.Kind)

	_, err := db.Client.GetAll(c, query, &entities)
	if err != nil {
		crud.logAndWriteError(c, err)
		return
	}

	c.JSON(http.StatusOK, entities)
}

func (crud *CRUD[T, T2, T3]) GetMulti(c *gin.Context) {
	keyStrs, ok := c.Params.Get("keys")
	if !ok {
		crud.logAndWriteError(c, ErrInvalidInput)
		return
	}

	var keys []*datastore.Key
	for _, keyStr := range strings.Split(keyStrs, ",") {
		key, err := datastore.DecodeKey(keyStr)
		if err != nil {
			crud.logAndWriteError(c, err)
			return
		}

		keys = append(keys, key)
	}

	fmt.Println(keys)
	entities := make([]T, len(keys))
	err := db.Client.GetMulti(c, keys, entities)
	if err != nil {
		crud.logAndWriteError(c, err)
		return
	}

	c.JSON(http.StatusOK, entities)
}
