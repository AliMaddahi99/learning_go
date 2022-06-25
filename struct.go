package main

import "fmt"

type person struct {
	name   string
	family string
	age    int
}

func (p person) nameFamilyAge(name, family string, age int) (string, string, int) {
	p.name = name
	p.family = family
	p.age = age
	return p.name, p.family, p.age
}

func main() {
	var ali person
	// var ali = person{name: "ali", family: "maddahi", age: 22}
	// fmt.Println(ali)

	fmt.Println(ali.nameFamilyAge("ali", "maddahi", 22))
}
