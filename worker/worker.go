package worker

import (
	"fmt"
	"sync"
	"time"
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

func NewGoroutinePool(workerSize int, items_per_worker int) *GoroutinePool {
	gp := &GoroutinePool{
		queue: make(chan work),
	}
	gp.AddWorkers(workerSize, items_per_worker)
	return gp
}

func (gp *GoroutinePool) Close() {
	close(gp.queue)
	gp.wg.Wait()
}

func (gp *GoroutinePool) ScheduleWork(fn WorkFunc) {
	gp.queue <- work{fn: fn}

}

func (gp *GoroutinePool) AddWorkers(numWorkers int, items_per_worker int) {
	gp.wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func(workerID int, gp *GoroutinePool) {

			defer gp.wg.Done()

			count := 0
			for job := range gp.queue {
				if items_per_worker == count {
					break
				}
				time.Sleep(time.Millisecond)
				job.fn.Run()
				time.Sleep(time.Millisecond)
				count++
			}

			fmt.Println(fmt.Sprintf("Worker %d executed %d tasks", workerID, count))
		}(i, gp)
	}
}
