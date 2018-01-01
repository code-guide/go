package main

import "fmt"

type Student struct {
	name string
	age  int
}

func SturctDemo() {
	var Tom Student
	Tom.name = "Tom"
	Tom.age = 20
	fmt.Println(Tom)

	students := []Student{
		{name: "Jim", age: 20},
	}
	fmt.Println(students)
}
