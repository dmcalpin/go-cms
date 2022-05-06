package crud

import (
	"net/http"

	"github.com/dmcalpin/go-cms/db"
	"github.com/gin-gonic/gin"
)

func (crud *CRUD[T, T2, T3]) Update(c *gin.Context) {
	key, err := crud.decodeKeyParam(&c.Params)
	if err != nil {
		crud.logAndWriteError(c, err)
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

	// Update the entity
	// Parse JSON
	input := new(T3)
	err = c.Bind(input)
	if err != nil {
		crud.logAndWriteError(c, err)
		return
	}

	e.Patch(input)

	client := db.Client
	updatedKey, err := client.Put(c, key, e)
	if err != nil {
		crud.logAndWriteError(c, err)
		return
	}

	updatedEntity := entity.New()
	err = db.Client.Get(c, updatedKey, updatedEntity)
	if err != nil {
		crud.logAndWriteError(c, err)
		return
	}

	c.JSON(http.StatusOK, updatedEntity)
}
