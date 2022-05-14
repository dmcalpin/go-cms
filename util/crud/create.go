package crud

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (crud *CRUD[T]) Create(c *gin.Context) {
	// _ = c.MustGet(gin.AuthUserKey).(string)
	var t T
	input := t.New(nil)

	err := c.Bind(input)
	if err != nil {
		crud.logAndWriteError(c, err)
		return
	}

	// optional hook to validate input
	err = input.Validate()
	if err != nil {
		crud.logAndWriteError(c, err)
		return
	}

	err = input.NewKey(nil, nil)
	if err != nil {
		crud.logAndWriteError(c, err)
		return
	}

	input.SetCreatedAt()
	input.SetUpdatedAt()

	err = input.SaveAndGet(c)
	if err != nil {
		crud.logAndWriteError(c, err)
		return
	}

	c.JSON(http.StatusCreated, input)
}
