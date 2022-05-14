package crud

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (crud *CRUD[T]) Delete(c *gin.Context) {
	key, err := crud.decodeKeyParam(&c.Params)
	if err != nil {
		crud.logAndWriteError(c, err)
		return
	}

	var t T
	t.New(key).Delete(c)
	if err != nil {
		crud.logAndWriteError(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
