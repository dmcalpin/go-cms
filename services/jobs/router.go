package jobs

import (
	"cloud.google.com/go/datastore"
	"github.com/gin-gonic/gin"

	"github.com/dmcalpin/go-cms/util/crud"
)

const JobKind = "Job"

type Job struct {
	Key         *datastore.Key `datastore:"__key__" json:"key"`
	Title       string         `datastore:"title" json:"title"`
	Description string         `datastore:"description" json:"description"`
}

func (u *Job) Patch(i interface{}) {
	input := i.(*JobUpdateInput)
	if input.Title != "" {
		u.Title = input.Title
	}
	if input.Description != "" {
		u.Description = input.Description
	}
}

func (j *Job) New() crud.Patchable {
	return &Job{}
}

type JobCreateInput struct {
	Title       string `datastore:"title" json:"title" binding:"required"`
	Description string `datastore:"description" json:"description" binding:"required"`
}

type JobUpdateInput struct {
	Title       string `datastore:"title" json:"title"`
	Description string `datastore:"description" json:"description"`
}

func AddRouter(r *gin.RouterGroup) {
	rg := r.Group("/jobs")

	jobCrud := crud.New[*Job, JobCreateInput, JobUpdateInput](JobKind)

	rg.POST("/", jobCrud.Create)
	rg.DELETE("/:key", jobCrud.Delete)
	rg.PUT("/:key", jobCrud.Update)
	rg.GET("/multi/:keys", jobCrud.GetMulti)
	rg.GET("/:key", jobCrud.Get)
	rg.GET("/", jobCrud.GetAll)
}
