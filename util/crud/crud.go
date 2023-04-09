package crud

import (
	"errors"
	"html/template"
	"net/http"

	"cloud.google.com/go/datastore"
	"github.com/dmcalpin/go-cms/db"
	"github.com/dmcalpin/go-cms/services/shared/templates"
	"github.com/gin-gonic/gin"
)

var ErrInvalidInput = errors.New("Invalid Input")

// T is the full GET model
// T2 is for the CREATE fields
// T3 is for the UPDATE fields
type CRUD[T db.Patchable] struct{}

func New[T db.Patchable]() *CRUD[T] {
	return &CRUD[T]{}
}

func (crud *CRUD[T]) logAndWriteError(c *gin.Context, err error) {
	c.Error(err)
	c.JSON(crud.errToHTTPError(err), c.Errors.JSON())
}

func (crud *CRUD[T]) logAndRenderError(c *gin.Context, err error) {
	c.Error(err)
	t := template.New("error.gohtml")
	templates.AddFuncs(t)
	t, parseErr := template.ParseFiles("services/shared/templates/layout_standard.gohtml", "services/shared/templates/error.gohtml")
	if parseErr != nil {
		panic(parseErr)
	}
	t.Execute(c.Writer, map[string]interface{}{
		"Status":  crud.errToHTTPError(err),
		"Message": err.Error(),
	})
}

func (crud *CRUD[T]) decodeKeyParam(params *gin.Params) (*datastore.Key, error) {
	keyParam := params.ByName("key")
	return datastore.DecodeKey(keyParam)
}

func (crud *CRUD[T]) errToHTTPError(err error) int {
	switch {
	case errors.Is(err, datastore.ErrNoSuchEntity):
		return http.StatusNotFound
	case errors.Is(err, datastore.ErrInvalidEntityType):
		return http.StatusNotAcceptable
	case errors.Is(err, datastore.ErrInvalidKey):
		return http.StatusNotAcceptable
	case errors.Is(err, ErrInvalidInput):
		return http.StatusBadRequest
	}

	return http.StatusInternalServerError
}
