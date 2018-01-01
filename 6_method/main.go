package main

import "fmt"

type Student struct {
	name string
	age  int
}

func (s Student) GetInfo() string {
	return fmt.Sprintf("name: %s, age: %d", s.name, s.age)
}

func main() {
	tom := Student{
		name: "Tom",
		age:  20,
	}

	fmt.Println(tom.GetInfo())
}
