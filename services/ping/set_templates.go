package ping

import (
	"html/template"
	"log"

	"github.com/gin-gonic/gin"
)

func SetTemplates(r *gin.Engine) {
	tmpl, err := template.ParseFiles("services/ping/ping.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	r.SetHTMLTemplate(tmpl)
}
