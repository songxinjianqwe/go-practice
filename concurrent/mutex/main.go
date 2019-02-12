package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex

func opAtomically() {
	mutex.Lock()
	defer mutex.Unlock()
	fmt.Println("opAtomically executing")
}

func main() {
	mutex.Lock()
	defer mutex.Unlock()
	fmt.Println("main executing")
	opAtomically()
}