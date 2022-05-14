package crud

import (
	"errors"
	"net/http"

	"cloud.google.com/go/datastore"
	"github.com/dmcalpin/go-cms/db"
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
