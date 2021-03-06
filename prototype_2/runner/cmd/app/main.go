package main

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/marcoshuck/runner/prototype_2/runner"
)

func main() {
	p := runner.NewPlatform()

	ctx := context.Background()
	dep := runner.NewDeployment(uuid.New())

	launchInstancesJob := runner.NewLaunchInstancesJob(p)
	waitInstancesJob := runner.NewWaitInstancesJob()
	removeInstancesJob := runner.NewRemoveInstancesJob(p)

	errorJob := runner.NewErrorJob(errors.New("some error"))

	pipeline := runner.NewPipeline(
		[]runner.Job{
			launchInstancesJob,
			waitInstancesJob,
			errorJob,
		},
		[]runner.Job{
			removeInstancesJob,
		},
	)

	pipeline.Run(ctx, dep)
}
