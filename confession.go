package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	LOVE = 2
)

type human struct {
	Name    string
	Feeling int
}

func (me *human) isLove(another *human) bool {
	return me.Feeling >= LOVE
}

func (me *human) CultivateFeeling(another *human) {
	fmt.Printf("逛街 with %s\n", another.Name)
	fmt.Printf("旅游 with %s\n", another.Name)
	fmt.Printf("吃吃吃 with %s\n", another.Name)
	fmt.Printf("玩玩玩 with %s\n", another.Name)
	fmt.Printf("牵绊++\n\n")
	me.Feeling++
	another.Feeling++
}

func (me *human) ConfessTo(another *human) bool {
	if another.isLove(me) {
		fmt.Printf("I love you, %s\n", another.Name)
		return true
	} else {
		fmt.Println("Game over")
		return false
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your name: ")
	scanner.Scan()
	yourName := scanner.Text()
	fmt.Print("Enter your lover's name: ")
	scanner.Scan()
	yourLoverName := scanner.Text()
	lover := &human{Name: yourLoverName}
	you := &human{Name: yourName}
	fmt.Print("你是否认为表白是最终胜利时的号角，而非发起冲锋的冲锋号？")
	scanner.Scan()
	answer := scanner.Text()
	if answer == "yes" {
		// right operation
		for !lover.isLove(you) {
			you.CultivateFeeling(lover)
		}
		if you.ConfessTo(lover) && lover.ConfessTo(you) {
			fmt.Printf("完结，撒花！\n")
		}
	} else {
		// wrong operation
		for !lover.isLove(you) {
			// endless loop
			you.ConfessTo(lover)
		}
	}
}
