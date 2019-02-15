package slice

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"testing"
)

type Person struct {
	Name string
	Age  int
	Cars []Car
}

type Car struct {
	Name string
}

func createPerson(Name string, Age int, carNames []string) *Person {
	person := Person{}
	person.Name = Name
	person.Age = Age
	var cars []Car
	for _, carName := range carNames {
		cars = append(cars, createCar(carName))
	}
	person.Cars = cars
	return &person
}

var encoder *gob.Encoder
var decoder *gob.Decoder

func init() {
	var buf bytes.Buffer
	encoder = gob.NewEncoder(&buf)
	decoder = gob.NewDecoder(&buf)
}

func createCar(name string) Car {
	car := Car{}
	car.Name = name
	return car
}

var allPersons = []*Person{
	createPerson("p1", 1, []string{"c1", "c2"}),
	createPerson("p2", 2, []string{"c1", "c2"}),
}

func (this *Person) Clone() (Person, error) {
	if err := encoder.Encode(*this); err != nil {
		return Person{}, err
	}
	var person Person
	err := decoder.Decode(person)
	if err != nil {
		return Person{}, err
	}
	return person, nil
}

func getPersonList() []Person {
	var personList []Person
	for _, p := range allPersons {
		pCopy := *p
		pCopy.Cars = make([]Car, len(p.Cars))
		copy(pCopy.Cars, p.Cars)
		personList = append(personList, pCopy)
	}
	return personList
}

func TestCopy(t *testing.T) {
	list := getPersonList()
	for _, p := range list {
		fmt.Printf("%#v\n", p)
	}
	allPersons[0].Name = "p3"
	allPersons[0].Cars[0].Name = "c3"
	for _, p := range list {
		fmt.Printf("%#v\n", p)
	}
}
