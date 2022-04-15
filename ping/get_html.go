package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getHtml(c *gin.Context) {

	data := map[string]interface{}{
		"Name": "dave",
	}

	c.HTML(http.StatusOK, "ping.gohtml", data)
}
