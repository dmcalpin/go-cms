package crud

import (
	"errors"

	"cloud.google.com/go/datastore"
	"github.com/dmcalpin/go-cms/db"
	"github.com/gin-gonic/gin"
)

func (crud *CRUD[T]) GetOneHTML(c *gin.Context) {
	var t T

	key, err := crud.decodeKeyParam(&c.Params)
	if err != nil {
		crud.logAndRenderError(c, errors.New("key param required"))
		return
	}

	entity := t.New(key)
	err = db.Get(c, entity)
	if err != nil {
		crud.logAndRenderError(c, err)
		return
	}

	err = entity.GetOneTemplate().Execute(c.Writer, entity)
	if err != nil {
		crud.logAndRenderError(c, err)
		return
	}
}

func (crud *CRUD[T]) GetAllHTML(c *gin.Context) {
	var t T
	kind := t.New(nil).GetKind()
	var entities []T

	query := datastore.NewQuery(kind)

	_, err := db.Client.GetAll(c, query, &entities)
	if err != nil {
		crud.logAndRenderError(c, err)
		return
	}

	err = t.GetMultiTemplate().Execute(c.Writer, entities)
	if err != nil {
		crud.logAndRenderError(c, err)
		return
	}
}
