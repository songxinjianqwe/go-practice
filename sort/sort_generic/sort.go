package main

import (
	"fmt"
	"sort"
	"strconv"
)

type Person struct {
	Name string
	Age  int
}
type PersonList []Person

func (pl PersonList) Len() int {
	return len(pl)
}

func (pl PersonList) Less(i, j int) bool {
	return pl[i].Age > pl[j].Age
}

func (pl PersonList) Swap(i, j int) {
	pl[i], pl[j] = pl[j], pl[i]
}

func NewPerson(name string, age int) Person {
	person := Person{}
	person.Name = name
	person.Age = age
	return person
}

func main() {
	var personList PersonList
	for i := 0; i < 10; i++ {
		personList = append(personList, NewPerson("P"+strconv.Itoa(i), i))
	}
	sort.Sort(personList)
	for _, p := range personList {
		fmt.Printf("%#v\n", p)
	}
}
