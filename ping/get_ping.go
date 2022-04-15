package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Ping struct {
	Message string `json:"message"`
}

func getPing(c *gin.Context) {

	c.JSON(http.StatusOK, Ping{"pong"})
}
