package worker

import (
	"fmt"
	"sync"
)

type WorkFunc interface {
	Run()
}

type GoroutinePool struct {
	queue chan work
	wg    sync.WaitGroup
}

type work struct {
	fn WorkFunc
}

func NewGoroutinePool(workerSize int) *GoroutinePool {
	gp := &GoroutinePool{
		queue: make(chan work),
	}
	gp.AddWorkers(workerSize)
	return gp
}

func (gp *GoroutinePool) Close() {
	close(gp.queue)
	gp.wg.Wait()
}

func (gp *GoroutinePool) ScheduleWork(fn WorkFunc) {
	gp.queue <- work{fn: fn}

}

func (gp *GoroutinePool) AddWorkers(numWorkers int) {
	gp.wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func(workerID int) {
			count := 0
			for job := range gp.queue {
				job.fn.Run()
				count++
			}

			fmt.Println(fmt.Sprintf("Worker %d executed %d tasks", workerID, count))
			gp.wg.Done()

		}(i)
	}
}
