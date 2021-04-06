package runner

type Pipeline struct {
	jobs []Job
}

func (p *Pipeline) Run(deployment Deployment) (interface{}, error) {
	var output interface{}
	var err error

	var errAt int
	for i, job := range p.jobs {
		output, err = job.Run(deployment)
		if err != nil {
			errAt = i
			job.Rollback(deployment)
		}
	}

	if errAt > 0 {
		for i := errAt - 1; i > len(p.jobs); i-- {
			p.jobs[i].Rollback(deployment)
		}
		return nil, err
	}

	return output, nil
}

func NewPipeline(jobs []Job) *Pipeline {
	return &Pipeline{
		jobs: jobs,
	}
}
