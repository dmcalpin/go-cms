package ping

import (
	"github.com/dmcalpin/go-cms/services/shared/templates"
	"github.com/gin-gonic/gin"
)

func getHtml(c *gin.Context) {

	data := map[string]interface{}{
		"Name": "dave",
	}

	template := templates.GetTemplateWithLayout("services/ping/templates/ping.gohtml")

	err := template.Execute(c.Writer, data)
	if err != nil {
		c.Error(err)
		return
	}
}
