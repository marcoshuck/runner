package runner

import "errors"

type Job struct {
	preHooks     []Func
	execute      Func
	postHooks    []Func
	errorHandler Func
}

func (j *Job) Rollback(deployment Deployment) error {
	panic("implement me")
}

func (j *Job) Run(r Deployment) (interface{}, error) {

	if j.execute == nil {
		return nil, errors.New("missing execute function")
	}

	var out interface{}
	var err error

	for _, h := range j.preHooks {
		out, err = h(r)
		r.Input = out
		r.Error = err
		if err != nil {
			if j.errorHandler != nil {
				j.errorHandler(r)
			}
			return out, err
		}
	}

	out, err = j.execute(r)
	r.Input = out
	r.Error = err
	if err != nil {
		if j.errorHandler != nil {
			j.errorHandler(r)
		}
		return out, err
	}

	for _, h := range j.postHooks {
		out, err = h(r)
		r.Input = out
		r.Error = err
		if err != nil {
			if j.errorHandler != nil {
				j.errorHandler(r)
			}
			return out, err
		}
	}

	return out, nil
}

func NewJob() *Job {
	return &Job{}
}

func (j *Job) WithPreHook(runner Func) *Job {
	j.preHooks = append(j.preHooks, runner)
	return j
}

func (j *Job) WithExecute(runner Func) *Job {
	j.execute = runner
	return j
}

func (j *Job) WithPostHook(runner Func) *Job {
	j.postHooks = append(j.postHooks, runner)
	return j
}

func (j *Job) WithErrorHandler(runner Func) *Job {
	j.errorHandler = runner
	return j
}
