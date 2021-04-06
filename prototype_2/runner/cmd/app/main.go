package main

import (
	"context"
	"github.com/google/uuid"
	"github.com/marcoshuck/runner/prototype_2/runner"
)

func main() {
	p := runner.NewPlatform()

	// launchInstancesJob := runner.NewLaunchInstancesJob(p)

	ctx := context.Background()
	dep := runner.NewDeployment(uuid.New())

	waitInstancesJob := runner.NewWaitInstancesJob()

	removeInstancesJob := runner.NewRemoveInstancesJob(p)

	pipeline := runner.NewPipeline(
		[]runner.Job{
			// launchInstancesJob,
			waitInstancesJob,
		},
		[]runner.Job{
			removeInstancesJob,
		},
	)

	pipeline.Run(ctx, dep)
}
