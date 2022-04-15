package users

import (
	"net/http"

	"github.com/dmcalpin/go-cms/db"
	"github.com/gin-gonic/gin"
)

func deleteUser(c *gin.Context) {
	userKey, err := getKeyParam(&c.Params)
	if err != nil {
		logAndWriteError(c, err)
		return
	}

	client := db.GetClient()
	err = client.Delete(c, userKey)
	if err != nil {
		logAndWriteError(c, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}
