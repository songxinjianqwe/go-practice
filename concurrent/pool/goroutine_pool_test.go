package pool

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
)

type PrintNameTask struct {
	Name string
}

func (this* PrintNameTask) Run() {
	fmt.Printf("Name:%s\n", this.Name)
}

func TestGoroutinePool(t *testing.T) {
	pool := NewGoroutinePool(10)
	taskNum := 100
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(taskNum)
	for i := 0; i < taskNum; i++  {
		go func(index int) {
			pool.Run(&PrintNameTask{Name: "N" + strconv.Itoa(index)})
			waitGroup.Done()
		}(i)
	}
	waitGroup.Wait()
	pool.ShutDown()
}