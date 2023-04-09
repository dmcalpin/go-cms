package templates

import (
	"html/template"

	"cloud.google.com/go/datastore"
)

func AddFuncs(t *template.Template) {
	t.Funcs(template.FuncMap{
		"EncodeKey": func(k interface{}) string {
			key := k.(*datastore.Key)
			if key == nil {
				return ""
			}
			return key.Encode()
		},
	})
}
