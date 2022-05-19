package users

import (
	"errors"
	"strings"

	"cloud.google.com/go/datastore"

	"github.com/dmcalpin/go-cms/db"
)

const UserKind = "User"

type User struct {
	db.DatastoreModel
	EmailAddress string         `datastore:"emailAddress" json:"emailAddress"`
	Password     string         `datastore:"password" json:"-"`
	Job          *datastore.Key `datastore:"job" json:"job"`
}

func (u *User) New(key *datastore.Key) db.Patchable {
	user := &User{}
	user.Kind = UserKind
	user.Key = key

	return user
}

func (u *User) Patch(i interface{}) {
	input := i.(*User)
	if input.EmailAddress != "" {
		u.EmailAddress = input.EmailAddress
	}
	if input.Password != "" {
		u.Password = input.Password
	}
	if input.Job != nil {
		u.Job = input.Job
	}
}

func (u *User) Validate() error {
	if !strings.Contains(u.EmailAddress, "@") {
		return errors.New("bad email address, must contain '@'")
	}
	return nil
}
