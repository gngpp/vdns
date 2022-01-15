package pool

import "sync"

// Task task struct
type Task struct {
	// task ID
	ID string
	// task function
	fn func(payloadType Payload)
	// task payload
	Payload Payload
}

// Worker worker struct
type Worker struct {
	// task ID
	ID int64
	// task queue
	taskQueue chan Task
	done      bool
	wait      bool
	wg        *sync.WaitGroup
}

func (_this *Worker) Run() {
	go func() {
		for {
			// shutdown
			if _this.done {
				return
			}
			task := <-_this.taskQueue
			task.fn(task.Payload)

			if _this.wait {
				_this.wg.Done()
			}
		}
	}()
}
