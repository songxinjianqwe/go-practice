package pool

import (
	"sync"
)

type Runnable interface {
	Run()
}

type GoroutinePool struct {
	workQueue chan Runnable
	waitGroup sync.WaitGroup
}

func NewGoroutinePool(maxPoolSize int) *GoroutinePool {
	pool := GoroutinePool{
		workQueue: make(chan Runnable),
	}
	pool.waitGroup.Add(maxPoolSize)
	for i := 0; i < maxPoolSize; i++ {
		go func() {
			for runnable := range pool.workQueue {
				runnable.Run()
			}
			pool.waitGroup.Done()
		}()
	}
	return &pool
}

func (this *GoroutinePool) Run(r Runnable) {
	this.workQueue <- r
}

func (this *GoroutinePool) ShutDown() {
	close(this.workQueue)
	this.waitGroup.Wait()
}
