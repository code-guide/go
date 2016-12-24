package main

import (
	"fmt"
)

func main() {
	setName("demo")
	fmt.Println(getName())

	name, age := getInfo()
	fmt.Println(name, age)

	fmt.Println(multi(1, 2, 3, 4))

	Change(age)
	fmt.Println(age)

	ChangeAge(&age)
	fmt.Println(age)
}

var name string
var age uint

func setName(str string) {
	name = str
}

// 声明函数
func getName() string {
	return name
}

// 多个返回值 => 返回值可命名
func getInfo() (name string, age uint) {
	return name, age
}

// 可变参数
func multi(nums ...int) int {
	result := 0
	for _, num := range nums {
		result += num
	}
	return result
}

// 按值传递
func Change(num uint) {
	num = num + 2
}

// 传入指针
func ChangeAge(num *uint) {
	*num = *num + 11
}
