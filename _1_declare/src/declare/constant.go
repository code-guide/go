/**
 * 声明常量
 */
package main

import (
	"fmt"
)

// 声明：const name = value

// 可以全局可以局部，可以直接赋值
const AGE = 11

// 可以先声明后赋值
const MALE = "male"

// 不可修改常量值
// MALE = "aaa"

// 不可重复声明
// const MALE = "ASDASD"

// 枚举的支持
const (
	a = iota
	b = iota
)

const c = iota

func main() {
	const NAME = "demo"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
