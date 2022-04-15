package users

import (
	"errors"
	"net/http"

	"cloud.google.com/go/datastore"
	"github.com/dmcalpin/go-cms/db"
	"github.com/gin-gonic/gin"
)

var ErrInvalidInput = errors.New("Invalid Input")

func getUser(c *gin.Context) {
	userKey, err := getKeyParam(&c.Params)
	if err != nil {
		logAndWriteError(c, ErrInvalidInput)
		return
	}

	user, err := getUserByKey(c, userKey)
	if err != nil {
		logAndWriteError(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func getAllUsers(c *gin.Context) {
	var users []*User

	client := db.GetClient()
	query := datastore.NewQuery(UserKind)

	_, err := client.GetAll(c, query, &users)
	if err != nil {
		logAndWriteError(c, err)
	}

	c.JSON(http.StatusOK, users)
}

func getKeyParam(params *gin.Params) (*datastore.Key, error) {
	userKeyParam := params.ByName("key")
	return datastore.DecodeKey(userKeyParam)

}

func getUserByKey(c *gin.Context, userKey *datastore.Key) (*User, error) {
	client := db.GetClient()

	user := new(User)
	err := client.Get(c, userKey, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func logAndWriteError(c *gin.Context, err error) {
	c.Error(err)
	c.JSON(datastoreToHTTPError(err), nil)
}
