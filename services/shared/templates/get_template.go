package templates

import (
	"html/template"
	"strings"
)

func GetTemplateWithLayout(filePath string) *template.Template {
	fileParts := strings.Split(filePath, "/")
	t := template.New(fileParts[len(fileParts)-1])

	AddFuncs(t)

	t, err := t.ParseFiles("services/shared/templates/layout_standard.gohtml", filePath)
	if err != nil {
		panic(err)
	}
	return t
}
