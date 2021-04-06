package runner

import (
	"context"
)

type Pipeline interface {
	Run(ctx context.Context, dep Deployment) Deployment
	Rollback(ctx context.Context, dep Deployment) Deployment
}

type pipeline struct {
	JobList          []Job
	RollbackHandlers []Job
	LastJobAt        int
	Error            error
}

func (p *pipeline) Run(ctx context.Context, dep Deployment) Deployment {
	for i, job := range p.JobList {

		dep = job(ctx, dep)

		if dep.HasError() {
			dep = p.Rollback(ctx, dep)
			return dep
		}

		p.LastJobAt = i
	}

	return dep
}

func (p *pipeline) Rollback(ctx context.Context, dep Deployment) Deployment {
	for i := len(p.RollbackHandlers) - 1; i >= 0; i-- {
		dep = p.RollbackHandlers[i](ctx, dep)
	}
	return dep
}

func NewPipeline(jobs []Job, rollbackHandlers []Job) Pipeline {
	return &pipeline{
		JobList:          jobs,
		RollbackHandlers: rollbackHandlers,
	}
}
