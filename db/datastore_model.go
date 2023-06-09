package db

import (
	"context"
	"errors"
	"html/template"
	"reflect"
	"time"

	"cloud.google.com/go/datastore"
)

type Patchable interface {
	Patch(interface{})
	New(*datastore.Key) Patchable
	NewKey(interface{}, *datastore.Key) error
	SetKey(*datastore.Key)
	GetKey() *datastore.Key
	GetKind() string
	SetUpdatedAt()
	SetCreatedAt()
	Validate() error
	CreateTemplate() *template.Template
	UpdateTemplate() *template.Template
	GetMultiTemplate() *template.Template
	GetOneTemplate() *template.Template
}

type DatastoreModel struct {
	Key       *datastore.Key `datastore:"__key__" json:"key"`
	Kind      string         `datastore:"-" json:"-"`
	CreatedAt time.Time      `datastore:"created_at" json:"createdAt"`
	UpdatedAt time.Time      `datastore:"updated_at" json:"updatedAt"`
}

func (dm *DatastoreModel) NewKey(idOrName interface{}, parent *datastore.Key) error {
	switch v := idOrName.(type) {
	case nil:
		dm.Key = datastore.IncompleteKey(dm.Kind, parent)
	case int64:
		dm.Key = datastore.IDKey(dm.Kind, reflect.ValueOf(v).Int(), parent)
	case string:
		dm.Key = datastore.NameKey(dm.Kind, v, parent)
	default:
		return errors.New("key value must be nil, int64, or string. Got: " + reflect.TypeOf(v).String())
	}

	return nil
}

func (dm *DatastoreModel) SetKey(key *datastore.Key) {
	dm.Key = key
}

func (dm *DatastoreModel) GetKey() *datastore.Key {
	return dm.Key
}

func (dm *DatastoreModel) GetKind() string {
	return dm.Kind
}

func (dm *DatastoreModel) Validate() error {
	return nil
}

func (dm *DatastoreModel) SetCreatedAt() {
	dm.CreatedAt = time.Now()
}

func (dm *DatastoreModel) SetUpdatedAt() {
	dm.UpdatedAt = time.Now()
}

func Get(c context.Context, p Patchable) error {
	return Client.Get(c, p.GetKey(), p)
}

func Delete(c context.Context, p Patchable) error {
	return Client.Delete(c, p.GetKey())
}

func Save(c context.Context, p Patchable) error {
	updatedKey, err := Client.Put(c, p.GetKey(), p)
	if err != nil {
		return err
	}

	p.SetKey(updatedKey)

	return nil
}

func SaveAndGet(c context.Context, p Patchable) error {
	err := Save(c, p)
	if err != nil {
		return err
	}

	return Get(c, p)
}
