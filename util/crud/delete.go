package crud

import (
	"net/http"

	"github.com/dmcalpin/go-cms/db"
	"github.com/gin-gonic/gin"
)

func (crud *CRUD[T, T2, T3]) Delete(c *gin.Context) {
	key, err := crud.decodeKeyParam(&c.Params)
	if err != nil {
		crud.logAndWriteError(c, err)
		return
	}

	err = db.Client.Delete(c, key)
	if err != nil {
		crud.logAndWriteError(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
