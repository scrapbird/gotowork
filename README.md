# gotowork
Provides a simple and easy to use worker queue for golang.

## Example
```go
package main

import (
	"fmt"
	"github.com/scrapbird/gotowork"
)

func main () {
	// create the workerqueue
	workerQueue := make(WorkerQueue, 5)

	// create and start the workers
	workers := make([]Worker, 5)
	for i := range workers {
		workers[i] = NewWorker(i, workerQueue)
		workers[i].Start()
	}

	// add some jobs
	for i := 0; i < 5; i++ {
		// get a free worker
		worker := <-workerQueue
		// give it some work
		worker <- func () {
			fmt.Println("Hello")
		}
	}

	// tell the workers to stop waiting for work
	for i := range workers {
		workers[i].Stop()
	}

	// wait for the workers to quit
	for i := range workers {
		workers[i].WaitForFinish()
	}
}
```
