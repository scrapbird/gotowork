/*
Package gotowork implements a simple worker pool.

First we create a WorkerQueue that can hold up to 5 workers

	workerQueue := make(gotowork.WorkerQueue, 5)

Next we create 5 workers and start each one

	for i := range workers {
		workers[i] = gotowork.NewWorker(i, workerQueue)
		workers[i].Start()
	}

Then we add some jobs to the queue

	for i := 0; i < 5; i++ {
		// get a free worker
		worker := <-workerQueue
		// give it some work (this can be any function)
		worker <- func () {
			fmt.Println("Hello")
		}
	}

Once all work has been dispatched we tell the workers to stop

	for i := range workers {
		workers[i].Stop()
	}

Then wait for all the workers to finish the current jobs

    for i := range workers {
        workers[i].WaitForFinish()
    }
*/
package gotowork
