package crud

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (crud *CRUD[T, T2, T3]) Delete(c *gin.Context) {
	key, err := crud.decodeKeyParam(&c.Params)
	if err != nil {
		crud.logAndWriteError(c, err)
		return
	}

	err = crud.DB.Delete(c, key)
	if err != nil {
		crud.logAndWriteError(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
