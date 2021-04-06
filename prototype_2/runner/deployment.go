package runner

import (
	"github.com/google/uuid"
)

type Deployment interface {
	Get(key string) interface{}
	Set(key string, data interface{})
	Exists(key string) bool
	Fail(err error)
	HasError() bool
	Error() error
}

func NewDeployment(uuid uuid.UUID) Deployment {
	return &deployment{
		UUID: uuid,
		data: make(map[string]interface{}),
	}
}

type deployment struct {
	data map[string]interface{}
	UUID uuid.UUID
	err  error
}

func (d *deployment) HasError() bool {
	return d.err != nil
}

func (d *deployment) Error() error {
	return d.err
}

func (d *deployment) Fail(err error) {
	d.err = err
}

func (d *deployment) Exists(key string) bool {
	_, ok := d.data[key]
	return ok
}

func (d *deployment) Get(key string) interface{} {
	return d.data[key]
}

func (d *deployment) Set(key string, data interface{}) {
	d.data[key] = data
}
