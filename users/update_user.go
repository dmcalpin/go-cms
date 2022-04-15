package users

import (
	"net/http"

	"github.com/dmcalpin/go-cms/db"
	"github.com/gin-gonic/gin"
)

type UserUpdateInput struct {
	EmailAddress string `json:"emailAddress" binding:"required"`
}

func updateUser(c *gin.Context) {
	userKey, err := getKeyParam(&c.Params)
	if err != nil {
		logAndWriteError(c, err)
		return
	}

	user, err := getUserByKey(c, userKey)
	if err != nil {
		logAndWriteError(c, err)
		return
	}

	// Update the user
	// Parse JSON
	userInput := new(UserUpdateInput)
	err = c.Bind(userInput)
	if err != nil {
		logAndWriteError(c, err)
		return
	}

	user.EmailAddress = userInput.EmailAddress

	client := db.GetClient()
	_, err = client.Put(c, user.Key, user)
	if err != nil {
		logAndWriteError(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}
