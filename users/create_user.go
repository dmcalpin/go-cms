package users

import (
	"errors"
	"net/http"

	"github.com/dmcalpin/go-cms/db"
	"github.com/gin-gonic/gin"

	"cloud.google.com/go/datastore"
)

const UserKind = "User"

type User struct {
	Key          *datastore.Key `datastore:"__key__"`
	EmailAddress string         `datastore:"emailAddress"`
	Password     string         `datastore:"password"`
}

type UserCreateInput struct {
	EmailAddress string `json:"emailAddress" binding:"required"`
	Password     string `json:"password" binding:"required"`
}

func createUser(c *gin.Context) {
	// _ = c.MustGet(gin.AuthUserKey).(string)

	// Parse JSON
	userInput := new(UserCreateInput)
	err := c.Bind(userInput)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	user := User{
		EmailAddress: userInput.EmailAddress,
		Password:     userInput.Password,
	}

	client := db.GetClient()
	userKey, err := client.Put(c, datastore.IncompleteKey(UserKind, nil), &user)
	if err != nil {
		c.Error(err)
		c.JSON(datastoreToHTTPError(err), nil)
		return
	}

	user.Key = userKey
	c.JSON(http.StatusCreated, &user)

}

func datastoreToHTTPError(datastoreErr error) int {
	if errors.Is(datastoreErr, datastore.ErrNoSuchEntity) {
		return http.StatusNotFound
	} else if errors.Is(datastoreErr, datastore.ErrInvalidEntityType) {
		return http.StatusNotAcceptable
	} else if errors.Is(datastoreErr, datastore.ErrInvalidKey) {
		return http.StatusNotAcceptable
	} else if errors.Is(datastoreErr, ErrInvalidInput) {
		return http.StatusNotAcceptable
	}

	return http.StatusInternalServerError
}
