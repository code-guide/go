package main

import "fmt"

type Animal interface {
	say() string
}

type Person struct {
	name string
	age  int
}

func (p Person) say() string {
	return fmt.Sprintf("name: %s, age: %d.", p.name, p.age)
}

func main() {
	tom := Person{name: "Tom", age: 20}

	var student Animal = tom
	fmt.Println(student.say())
}
