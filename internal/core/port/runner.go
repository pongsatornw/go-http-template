package port

// Runner defines a generic lifecycle interface for any long-running component,
// such as servers, background workers, or schedulers.
//
// It provides methods to start and stop the component in a controlled manner.
type Runner interface {
	// Start begins the execution of the component.
	// It should perform initialization and start any required routines.
	// If an error occurs during startup, it should be sent to startErrCh.
	Start(startErrCh chan<- error)

	// Shutdown gracefully stops the component and cleans up any resources.
	// It should block until the shutdown is complete.
	//
	// Expect external call to this function when shutdown process. eg. defer
	Shutdown()
}
