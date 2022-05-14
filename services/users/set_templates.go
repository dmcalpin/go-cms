package users

import (
	"html/template"
	"log"

	"github.com/gin-gonic/gin"
)

func SetTemplates(r *gin.Engine) {
	tmpl, err := template.ParseFiles("services/users/templates/users_list.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	r.SetHTMLTemplate(tmpl)
}
