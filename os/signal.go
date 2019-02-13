package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

/**
如果当前进程收到了未自定义处理方法的信号，就会执行由OS指定的默认操作。
如果收到了自定义处理方法的信号，就会放到对应的channel中，由我们自己来处理这个信号，仅仅打印一些信息是不够的。
特例：SIGKILL和SIGSTOP是无法自定义处理方法的，只会执行OS的默认操作，根本原因是它们向OS的超级用户提供了使进程停止或终止的可靠方法
*/
func registerAndListenerThenStopSignals() {
	receivedSignals := make(chan os.Signal, 1)
	customProcessingSignals := []os.Signal{syscall.SIGINT, syscall.SIGQUIT}
	fmt.Printf("Set notification for %s... [customProcessingSignals]\n", customProcessingSignals)
	// 类似于addListener，当当前进程真正收到信号时，会放入到channel中，我们可以使用一个goroutine来监听这个channel
	signal.Notify(receivedSignals, customProcessingSignals...)
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(1)
	go func() {
		for sig := range receivedSignals {
			fmt.Printf("received %s\n", sig)
		}
		waitGroup.Done()
	}()
	// 取消掉在之前调用signal.Notify函数时告知signal处理程序需要自行处理若干信号的行为。
	signal.Stop(receivedSignals)
	close(receivedSignals)
	waitGroup.Done()
}

func main() {
	registerAndListenerThenStopSignals()
}
