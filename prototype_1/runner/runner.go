package runner

import "context"

type Deployment struct {
	ctx   context.Context
	Input interface{}
	Error error
}

type Func func(deployment Deployment) (interface{}, error)

type RollbackFunc func(deployment Deployment) error
