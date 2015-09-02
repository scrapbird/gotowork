// Provides a simple and easy to use worker queue.
package gotowork

// Define work as a first class function
type Work func()

// Define the worker queue
type WorkerQueue chan Inbox

// Define the inbox
type Inbox chan Work

// This is the worker type, this is used to keep track of which queue the worker is in and the worker id, signaling channels etc.
type Worker struct {
	id          int
	inbox       Inbox
	workerQueue WorkerQueue
	quit        chan bool
	done        chan bool
}

// Creates a new worker struct and initializes it.
//
//  workerQueue := make(WorkerQueue, 5) // Create worker queue that fits up to 5 workers.
//  worker := NewWorker(1, workerQueue) // Create a new worker in the queue.
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

// Tells the worker to start subscribing to the worker queue and executing jobs.
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

// Tells the worker to stop once it has finished it's current job.
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

// Waits for a worker to finish before returning.
func (w Worker) WaitForFinish() {
	<-w.done
}
