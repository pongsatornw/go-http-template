package port

type Server interface {
	// startErrCh: send error to startErrCh if server failed to start
	Start(startErrCh chan<- error)
	Shutdown()
}
