package crud

import (
	"errors"
	"net/http"

	"cloud.google.com/go/datastore"
	"github.com/gin-gonic/gin"
)

type Patchable interface {
	Patch(interface{})
	New() Patchable
}

var ErrInvalidInput = errors.New("Invalid Input")

// T is the full GET model
// T2 is for the CREATE fields
// T3 is for the UPDATE fields
type CRUD[T Patchable, T2 any, T3 any] struct {
	// datastore document kind
	Kind string
}

func New[T Patchable, T2 any, T3 any](kind string) *CRUD[T, T2, T3] {
	return &CRUD[T, T2, T3]{
		Kind: kind,
	}
}

func (crud *CRUD[T, T2, T3]) logAndWriteError(c *gin.Context, err error) {
	c.Error(err)
	c.JSON(crud.errToHTTPError(err), nil)
}

func (crud *CRUD[T, T2, T3]) decodeKeyParam(params *gin.Params) (*datastore.Key, error) {
	keyParam := params.ByName("key")
	return datastore.DecodeKey(keyParam)
}

func (crud *CRUD[T, T2, T3]) errToHTTPError(err error) int {
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
