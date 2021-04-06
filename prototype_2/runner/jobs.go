package runner

import (
	"context"
	"errors"
	"fmt"
)

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
