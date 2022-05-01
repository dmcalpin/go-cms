package crud

import (
	"net/http"
	"reflect"

	"github.com/dmcalpin/go-cms/db"
	"github.com/gin-gonic/gin"
)

func (crud *CRUD[T, T2, T3]) Update(c *gin.Context) {
	key, err := crud.decodeKeyParam(&c.Params)
	if err != nil {
		crud.logAndWriteError(c, err)
		return
	}

	entity := new(T)
	err = crud.DB.Get(c, key, entity)
	if err != nil {
		crud.logAndWriteError(c, err)
		return
	}

	// Update the entity
	// Parse JSON
	input := new(T3)
	err = c.Bind(input)
	if err != nil {
		crud.logAndWriteError(c, err)
		return
	}

	entity = crud.patchValues(entity, input)

	client := db.GetClient()
	updatedKey, err := client.Put(c, key, entity)
	if err != nil {
		crud.logAndWriteError(c, err)
		return
	}

	updatedEntity := new(T)
	err = crud.DB.Get(c, updatedKey, updatedEntity)
	if err != nil {
		crud.logAndWriteError(c, err)
		return
	}

	c.JSON(http.StatusOK, updatedEntity)
}

func (crud *CRUD[T, T2, T3]) patchValues(a *T, b *T3) *T {
	aObj := reflect.ValueOf(a).Elem()
	bObj := reflect.ValueOf(b).Elem()
	bType := reflect.TypeOf(*b)
	for _, field := range reflect.VisibleFields(bType) {
		fieldName := field.Name
		bFieldVal := bObj.FieldByName(fieldName)
		if !bFieldVal.IsZero() {
			aObj.FieldByName(fieldName).Set(reflect.ValueOf(bFieldVal.Interface()))
		}

	}
	return a
}
