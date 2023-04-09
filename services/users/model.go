package users

import (
	"errors"
	"html/template"
	"strings"

	"cloud.google.com/go/datastore"

	"github.com/dmcalpin/go-cms/db"
	"github.com/dmcalpin/go-cms/services/shared/templates"
)

const UserKind = "User"

var createTemplate, multiTemplate, oneTemplate *template.Template

func init() {
	createTemplate = templates.GetTemplateWithLayout("services/users/templates/users_create.gohtml")
	multiTemplate = templates.GetTemplateWithLayout("services/users/templates/users_list.gohtml")
	oneTemplate = templates.GetTemplateWithLayout("services/users/templates/users_details.gohtml")
}

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

// Template config
func (u *User) CreateTemplate() *template.Template {
	return createTemplate
}
func (u *User) UpdateTemplate() *template.Template {
	return nil
}
func (u *User) GetMultiTemplate() *template.Template {
	return multiTemplate
}
func (u *User) GetOneTemplate() *template.Template {
	return oneTemplate
}
