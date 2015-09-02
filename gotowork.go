// Provides a simple and easy to use worker queue.
package gotowork

// Define work as a first class function
type Work func()

// Define the worker queue
type WorkerQueue chan chan Work

// Define the inbox
type Inbox chan Work

// Define the worker
type Worker struct {
	id          int         // id of the worker
	inbox       Inbox       // channel to send work requests to this worker
	workerQueue WorkerQueue // channel to register with to receive jobs
	quit        chan bool   // worker will quit if a message is received here
	done        chan bool   // used to signal that worker is finished
}

// Creates a worker
func NewWorker(id int, workerQueue WorkerQueue) Worker {
	worker := Worker{
		id:          id,
		inbox:       make(chan Work),
		workerQueue: workerQueue,
		quit:        make(chan bool),
		done:        make(chan bool),
	}
	return worker
}

// Starts the worker
func (w Worker) Start() {
	go func() {
		for {
			// signal that we are ready for work
			w.workerQueue <- w.inbox
			select {
			case work := <-w.inbox:
				work()
			case <-w.quit:
				w.done <- true
				return
			}
		}
	}()
}

// Requests the worker to stop
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

// Waits for a worker to finish
func (w Worker) WaitForFinish() {
	<-w.done
}
