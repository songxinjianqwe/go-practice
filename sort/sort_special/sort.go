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

func NewPerson(name string, age int) Person {
	person := Person{}
	person.Name = name
	person.Age = age
	return person
}

func main() {
	var personList []Person
	for i := 0; i < 10; i++ {
		personList = append(personList, NewPerson("P"+strconv.Itoa(i), i))
	}
	sort.Slice(personList, func(i, j int) bool {
		return personList[i].Age > personList[j].Age
	})
	for _, p := range personList {
		fmt.Printf("%#v\n", p)
	}
}
