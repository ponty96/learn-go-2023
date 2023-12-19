package main

import "fmt"


type person struct {
	name string
	age int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func updatePerson(p *person, age int) {
	p.age = age
}

func main() {
	wrong_age := 26

	ayomide := person{name: "Ayomide", age: wrong_age}

	fmt.Println("ayomide: ", ayomide)

	updatePerson(&ayomide, 27)

	fmt.Println("ayomide: ", ayomide)

	fmt.Println("pointer ayomide: ", &ayomide)

	inlineStruct := struct {
		name string
		level int
	}{
		"inlineStruct",
		1,
	}

	fmt.Println("inlineStruct: ", inlineStruct)
}
