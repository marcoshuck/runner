package runner

// Runner executes a Job.
type Runner interface {
	Run(deployment Deployment) (interface{}, error)
}

// RollbackHandler revert changes made by a Job.
type RollbackHandler interface {
	Rollback(deployment Deployment) error
}

// Setup
type Setup interface {
	SetupPreHook(f ...Func)
	SetupExecute(f Func)
	SetupPostHook(f ...Func)
	SetupErrorHandler(f Func)
}
