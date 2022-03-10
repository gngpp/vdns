package pool

import (
	"sync"
	"time"
	"vdns/lib/util/rand"
	"vdns/lib/vlog"
)

var log = vlog.Default()

// SpeedInterval default average speed interval
var SpeedInterval = time.Duration(10)

// Payload type (task type)
type Payload string

// --------------------------------------- //

// NewPool new Pool
func NewPool(maxWorkers int32, taskQueueLen int32, feedback bool) *Pool {
	taskQueue := make(chan Task, taskQueueLen)
	workers := make([]*Worker, taskQueueLen)

	if maxWorkers > taskQueueLen {
		log.Fatal("maxWorkers must smaller than the task taskQueueLen")
	}

	var wg sync.WaitGroup

	pool := &Pool{
		maxWorkers:   maxWorkers,
		taskQueueLen: taskQueueLen,
		taskQueue:    taskQueue,
		wait:         true,
		callback:     feedback,
		wg:           &wg,
		workers:      workers,
		doneTaskList: 0,
	}
	// start worker
	pool.Start()
	// start monitor
	pool.Monitor()

	return pool
}

type Pool struct {
	maxWorkers   int32
	doneTaskList int32
	taskQueue    chan Task
	taskQueueLen int32
	workers      []*Worker
	wg           *sync.WaitGroup
	wait         bool
	callback     bool
}

func (_this *Pool) AddTask(fn func(payload Payload), payload Payload) {
	uuid := rand.NewUUID()
	task := Task{
		ID:      uuid.ToString(),
		fn:      fn,
		Payload: payload,
	}
	if _this.wait {
		_this.wg.Add(1)
	}
	// bad
	_this.doneTaskList++
	_this.taskQueue <- task
}

func (_this *Pool) Monitor() {
	var lastDone int32
	// bad
	lastDone = _this.taskQueueLen
	var speed int32
	interval := time.NewTicker(SpeedInterval * time.Second)

	// real-time speed
	var realtimeLastDone int32
	// bad
	realtimeLastDone = _this.taskQueueLen
	var realtimeSpeed int32

	// interval monitor
	ticker := time.NewTicker(1 * time.Second)
	quit := make(chan struct{})

	// time cost
	tStart := time.Now()
	var costDuration time.Duration

	var callbackDoubleCheck = false

	initWorkers := _this.maxWorkers

	go func() {
		for {
			select {

			case <-ticker.C:
				tCurrent := time.Now()
				costDuration = tCurrent.Sub(tStart)

				// real-time speed
				tmpRealtimeSpeed := _this.doneTaskList - realtimeLastDone

				realtimeSpeed = tmpRealtimeSpeed
				realtimeLastDone = _this.doneTaskList

				log.Debugf("Start at: %s, time cost: %s, average speed: %d, read-time speed: %d, current workers: %d, done: %d     ", tStart.Format("15:04:05.000"), costDuration.String(), speed, realtimeSpeed, _this.maxWorkers, _this.doneTaskList)

			// callback mechanism!
			case <-interval.C:
				// average speed
				tmpSpeed := (_this.doneTaskList - lastDone) / int32(SpeedInterval)

				if _this.callback {
					var callbackMaxWorkers int32

					callbackMaxWorkers = _this.maxWorkers
					if tmpSpeed < speed {
						if callbackDoubleCheck == true {
							// callbackMaxWorkers = _this.maxWorkers - _this.maxWorkers*(speed-tmpSpeed)/speed
							callbackMaxWorkers -= callbackMaxWorkers / 10
							log.Infof("DoubleCheck Feedback \n")
							callbackDoubleCheck = false
						} else {
							callbackDoubleCheck = true
						}
					} else {
						if speed > 0 {
							// callbackMaxWorkers = _this.maxWorkers + _this.maxWorkers*(tmpSpeed-speed)/speed
							callbackMaxWorkers += callbackMaxWorkers / 10

						}
						callbackDoubleCheck = false
					}

					if callbackMaxWorkers < initWorkers {
						callbackMaxWorkers = initWorkers
					}
					if callbackMaxWorkers > _this.taskQueueLen {
						callbackMaxWorkers = _this.taskQueueLen
					}
					if callbackMaxWorkers > 2*_this.maxWorkers {
						callbackMaxWorkers = 2 * _this.maxWorkers
					}
					// callback worker numbers

					log.Debugf("Start at: %s, time cost: %s, average speed: %d, read-time speed: %d, current workers: %d, done: %d     ", tStart.Format("15:04:05.000"), costDuration.String(), speed, realtimeSpeed, callbackMaxWorkers, _this.doneTaskList)
					_this.callbackWorkers(callbackMaxWorkers)
				}

				speed = tmpSpeed
				lastDone = _this.doneTaskList

			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func (_this *Pool) MapRun(fn func(payload Payload), payloads []Payload) {
	// maybe reflect
	// http://blog.burntsushi.net/type-parametric-functions-golang/
	for _, payload := range payloads {
		_this.AddTask(fn, payload)
	}
}

func (_this *Pool) MapRunChan(fn func(payload Payload), fnfetch func() Payload) {
	for {
		payload := fnfetch()
		_this.AddTask(fn, payload)
	}
}

// Start coroutine pool
func (_this *Pool) Start() {
	if _this.taskQueueLen < _this.maxWorkers {
		log.Fatal("maxWorkers must smaller than the task taskQueueLen")
	}
	for i := int32(0); i < _this.maxWorkers; i++ {
		worker := &Worker{
			ID:        int64(i),
			taskQueue: _this.taskQueue,
			wait:      _this.wait,
			done:      false,
			wg:        _this.wg,
		}
		worker.Run()
		_this.workers[i] = worker
	}
}

func (_this *Pool) callbackWorkers(callbackMaxWorkers int32) {
	if callbackMaxWorkers > _this.maxWorkers {
		for i := _this.maxWorkers; i < callbackMaxWorkers; i++ {
			worker := &Worker{
				ID:        int64(i),
				taskQueue: _this.taskQueue,
				wait:      _this.wait,
				done:      false,
				wg:        _this.wg,
			}
			worker.Run()
			_this.workers[i] = worker
		}
	} else {
		for i := callbackMaxWorkers; i < _this.maxWorkers; i++ {
			_this.workers[i].done = true
		}
	}

	_this.maxWorkers = callbackMaxWorkers
}

func (_this *Pool) Wait() {
	_this.wg.Wait()
}
