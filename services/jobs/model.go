package jobs

import (
	"cloud.google.com/go/datastore"

	"github.com/dmcalpin/go-cms/db"
)

const JobKind = "Job"

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
