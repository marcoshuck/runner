package runner

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type Job func(ctx context.Context, dep Deployment) Deployment

func NewLaunchInstancesJob(p Platform) Job {
	return func(ctx context.Context, dep Deployment) Deployment {
		dep.Set("test", p.LaunchInstance())
		return dep
	}
}

var ErrNoInstancesAvailable = errors.New("no instances available")

func NewWaitInstancesJob() Job {
	return func(ctx context.Context, dep Deployment) Deployment {
		if !dep.Exists("test") {
			dep.Fail(ErrNoInstancesAvailable)
			return dep
		}

		instances := dep.Get("test").(int)
		fmt.Printf("Waiting for [%d] instances\n", instances)

		time.Sleep(5 * time.Second)

		return dep
	}
}

func NewRemoveInstancesJob(p Platform) Job {
	return func(ctx context.Context, dep Deployment) Deployment {
		if !dep.Exists("test") {
			return dep
		}

		if err := p.RemoveInstance(); err != nil {
			return dep
		}

		return dep
	}
}

func NewErrorJob(err error) Job {
	return func(ctx context.Context, dep Deployment) Deployment {
		dep.Fail(err)
		return dep
	}
}
