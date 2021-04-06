package main

import (
	"context"
	"github.com/google/uuid"
	"github.com/marcoshuck/runner/prototype_2/runner"
	"time"
)

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	dep := runner.NewDeployment(uuid.New())

	p := runner.NewPlatform()
	launchInstancesJob := runner.NewLaunchInstancesJob(p)
	waitInstancesJob := runner.NewWaitInstancesJob()
	removeInstancesJob := runner.NewRemoveInstancesJob(p)

	pipeline := runner.NewPipeline(
		[]runner.Job{
			launchInstancesJob,
			waitInstancesJob,
		},
		[]runner.Job{
			removeInstancesJob,
		},
	)

	pipeline.Run(ctx, dep)
}
