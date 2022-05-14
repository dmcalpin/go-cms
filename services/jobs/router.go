package jobs

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/gin-gonic/gin"

	"github.com/dmcalpin/go-cms/db"
	"github.com/dmcalpin/go-cms/util/crud"
)

const JobKind = "Job"

type Job struct {
	db.DatastoreModel
	Title       string `datastore:"title" json:"title"`
	Description string `datastore:"description" json:"description"`
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

func (j *Job) Get(c context.Context) error {
	return db.Client.Get(c, j.Key, j)
}

func (j *Job) Save(c context.Context) error {
	updatedKey, err := db.Client.Put(c, j.Key, j)
	if err != nil {
		return err
	}

	j.Key = updatedKey

	return nil
}

func (j *Job) SaveAndGet(c context.Context) error {
	err := j.Save(c)
	if err != nil {
		return err
	}

	return j.Get(c)
}

func (j *Job) Delete(c context.Context) error {
	return db.Client.Delete(c, j.Key)
}

func (j *Job) New(key *datastore.Key) db.Patchable {
	job := &Job{}
	job.Kind = JobKind
	job.Key = key

	return job
}

func AddRouter(r *gin.RouterGroup) {
	rg := r.Group("/api/jobs")

	jobCrud := crud.New[*Job]()

	rg.POST("/", jobCrud.Create)
	rg.DELETE("/:key", jobCrud.Delete)
	rg.PUT("/:key", jobCrud.Update)
	rg.GET("/multi/:keys", jobCrud.GetMulti)
	rg.GET("/:key", jobCrud.Get)
	rg.GET("/", jobCrud.GetAll)
}
