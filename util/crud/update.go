package crud

import (
	"net/http"

	"github.com/dmcalpin/go-cms/db"
	"github.com/gin-gonic/gin"
)

func (crud *CRUD[T]) Update(c *gin.Context) {
	key, err := crud.decodeKeyParam(&c.Params)
	if err != nil {
		crud.logAndWriteError(c, err)
		return
	}

	var t T
	entity := t.New(key)
	err = db.Get(c, entity)
	if err != nil {
		crud.logAndWriteError(c, err)
		return
	}

	// Update the entity
	// Parse JSON
	input := t.New(nil)
	err = c.Bind(input)
	if err != nil {
		crud.logAndWriteError(c, err)
		return
	}

	entity.Patch(input)

	err = entity.Validate()
	if err != nil {
		crud.logAndWriteError(c, err)
		return
	}

	entity.SetUpdatedAt()

	err = db.SaveAndGet(c, entity)
	if err != nil {
		crud.logAndWriteError(c, err)
		return
	}

	c.JSON(http.StatusOK, entity)
}
