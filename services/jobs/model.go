package jobs

import (
	"html/template"

	"cloud.google.com/go/datastore"

	"github.com/dmcalpin/go-cms/db"
	"github.com/dmcalpin/go-cms/services/shared/templates"
)

const JobKind = "Job"

var multiTemplate, oneTemplate *template.Template

func init() {
	multiTemplate = templates.GetTemplateWithLayout("services/jobs/templates/jobs_list.gohtml")
	oneTemplate = templates.GetTemplateWithLayout("services/jobs/templates/jobs_details.gohtml")
}

type Job struct {
	db.DatastoreModel
	Title       string `datastore:"title" json:"title"`
	Description string `datastore:"description" json:"description"`
}

func (j *Job) New(key *datastore.Key) db.Patchable {
	job := &Job{}
	job.Kind = JobKind
	job.Key = key

	return job
}

func (u *Job) Patch(i interface{}) {
	input := i.(*Job)
	if input.Title != "" {
		u.Title = input.Title
	}
	if input.Description != "" {
		u.Description = input.Description
	}
}

// Template config
func (u *Job) CreateTemplate() *template.Template {
	return nil
}
func (u *Job) UpdateTemplate() *template.Template {
	return nil
}
func (u *Job) GetMultiTemplate() *template.Template {
	return multiTemplate
}
func (u *Job) GetOneTemplate() *template.Template {
	return oneTemplate
}
