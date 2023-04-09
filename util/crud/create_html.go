package crud

import (
	"github.com/gin-gonic/gin"
)

func (crud *CRUD[T]) CreateHTML(c *gin.Context) {
	var t T
	entity := t.New(nil)

	err := entity.CreateTemplate().Execute(c.Writer, nil)
	if err != nil {
		crud.logAndRenderError(c, err)
		return
	}
}
