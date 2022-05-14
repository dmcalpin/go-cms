package crud

import (
	"net/http"
	"strings"

	"cloud.google.com/go/datastore"
	"github.com/dmcalpin/go-cms/db"
	"github.com/gin-gonic/gin"
)

func (crud *CRUD[T]) Get(c *gin.Context) {
	var t T

	key, err := crud.decodeKeyParam(&c.Params)
	if err != nil {
		crud.logAndWriteError(c, ErrInvalidInput)
		return
	}

	entity := t.New(key)
	err = entity.Get(c)
	if err != nil {
		crud.logAndWriteError(c, err)
		return
	}

	c.JSON(http.StatusOK, entity)
}

func (crud *CRUD[T]) GetAll(c *gin.Context) {
	var t T
	kind := t.New(nil).GetKind()
	var entities []T

	query := datastore.NewQuery(kind)

	_, err := db.Client.GetAll(c, query, &entities)
	if err != nil {
		crud.logAndWriteError(c, err)
		return
	}

	c.JSON(http.StatusOK, entities)
}

func (crud *CRUD[T]) GetMulti(c *gin.Context) {
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

	entities := make([]T, len(keys))
	err := db.Client.GetMulti(c, keys, entities)
	if err != nil {
		crud.logAndWriteError(c, err)
		return
	}

	c.JSON(http.StatusOK, entities)
}
