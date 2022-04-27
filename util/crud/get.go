package crud

import (
	"net/http"

	"cloud.google.com/go/datastore"
	"github.com/gin-gonic/gin"
)

func (crud *CRUD[T, T2, T3]) Get(c *gin.Context) {
	key, err := crud.decodeKeyParam(&c.Params)
	if err != nil {
		crud.logAndWriteError(c, ErrInvalidInput)
		return
	}

	entity := new(T)
	err = crud.DB.Get(c, key, entity)
	if err != nil {
		crud.logAndWriteError(c, err)
		return
	}

	c.JSON(http.StatusOK, entity)
}

func (crud *CRUD[T, T2, T3]) GetAll(c *gin.Context) {
	var entities []*T

	query := datastore.NewQuery(crud.Kind)

	_, err := crud.DB.GetAll(c, query, &entities)
	if err != nil {
		crud.logAndWriteError(c, err)
		return
	}

	c.JSON(http.StatusOK, entities)
}
