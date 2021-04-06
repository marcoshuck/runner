package main

import (
	"fmt"
	"github.com/marcoshuck/runner/prototype_1/runner"
)

func main() {
	job := runner.NewJob()

	job.WithPreHook(print).WithExecute(print).WithPostHook(print).WithErrorHandler(print)

	output, err := job.Run(runner.Deployment{
		Input: "Some test",
		Error: nil,
	})

	fmt.Println("Output:", output)
	fmt.Println("Error:", err)
}

func print(r runner.Deployment) (interface{}, error) {
	fmt.Println("Printing request", r.Input)

	r.Input = fmt.Sprintf("%v - printed", r.Input)

	return r.Input, nil
}
