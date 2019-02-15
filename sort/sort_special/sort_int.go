package main

import (
	"fmt"
	"sort"
)

func main() {
	n := 10
	intArr := make([]int, n)
	for i := 0; i < n; i++ {
		intArr[i] = n - i - 1
	}
	fmt.Printf("%#v\n", intArr)
	sort.Ints(intArr)
	fmt.Printf("%#v\n", intArr)
}
