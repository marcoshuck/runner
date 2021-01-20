package runner

type emptyRunner struct{}

func (e emptyRunner) Run(deployment Deployment) (interface{}, error) {
	return nil, nil
}

func Empty() Runner {
	return &emptyRunner{}
}
