package product

/*
*
- make worker pool to handle concurrent updates
- every product update request gets added into product channel(FIFO queue)
- users can verify updates status with get product details call
*/
type WorkerPool interface {
	// Start gets the workerpool ready to process jobs, and should only be called once
	Start()
	// Stop stops the workerpool, tears down any required resources,
	// and should only be called once
	Stop()
	// AddWork adds a task for the worker pool to process. It is only valid after
	// Start() has been called and before Stop() has been called.
	AddWork(Task)
}

type Task interface {
	// Execute performs the work
	Execute() error
	// OnFailure handles any error returned from Execute()
	OnFailure(error)
}

func Init() {
	ch := make(chan Product, 2)
	p := Product{}
	ch <- p
}
