package runner

// Runner executes a Job.
type Runner interface {
	Run(deployment Deployment) (interface{}, error)
}

// RollbackHandler revert changes made by a Job.
type RollbackHandler interface {
	Rollback(deployment Deployment) error
}

// Setup groups a set of methods to set up hooks, error handlers and execute functions.
// Not implemented yet.
type Setup interface {
	SetupPreHook(f ...Func)
	SetupExecute(f Func)
	SetupPostHook(f ...Func)
	SetupErrorHandler(f Func)
}
