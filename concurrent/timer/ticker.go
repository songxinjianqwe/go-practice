package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(2 * time.Second)
	// ticker是每隔2s在channel里放一个元素，类似于scheduleAtFixedRate
	for range ticker.C {
		fmt.Printf("before execution: %s\n", time.Now())
		time.Sleep(3 * time.Second)
		fmt.Printf("after execution: %s\n", time.Now())
	}
}
