package main

import "fmt"

// name 全局变量
var name = "demo"

// PI 常量
const PI = 3.14

func main() {

	// 局部变量简短声明
	localName := "local name"

	fmt.Println(name)
	fmt.Println(PI)
	fmt.Println(localName)
}
