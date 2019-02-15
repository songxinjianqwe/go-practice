package main

import (
	"fmt"
	"sync"
)

func printConcurrently(words []string, waitGroup *sync.WaitGroup) {
	for _, word := range words {
		go func(aWord string) {
			waitGroup.Add(1)
			defer waitGroup.Done()
			fmt.Println(aWord)
		}(word)
	}
}

/**
main函数如果要作为应用启动函数，则必须放在main包中
一个项目（项目并不是go中的概念）中，可以存在多个同名包
*/
func main() {
	waitGroup := sync.WaitGroup{}
	printConcurrently([]string{"Eric", "Harry", "Robert", "Jim", "Mark"}, &waitGroup)
	waitGroup.Wait()
	fmt.Println("finished!")
}
