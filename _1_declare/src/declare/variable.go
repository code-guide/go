/**
 * 声明变量
 */
package main

import (
	"fmt"
)

// 方法1 : var variableName type
// 可以定义在全局和局部，全局变量可以不使用
var str string

// 可以定义多个变量，可以声明并且初始化，可以同时声明并初始化多个变量
var a, b, c int = 1, 2, 3

// 函数外不可以进行赋值操作
// a = 2

// 可以合并声明
var (
	aa int    = 2
	bb string = "asdas"
)

// 支持类型推导
var d, e = 2, 3

func main() {

	// 方法1：局部变量必须使用
	var str string = "asdasd"
	fmt.Println(str)

	// 方法2：简短定义  variableName := value
	// 只能在局部定义，并且必须使用
	name := "demo"
	fmt.Println(name)

}
