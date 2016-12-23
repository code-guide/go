/**
 * 内置基本类型
 */
package main

import "fmt"
import "errors"

// ============= 布尔值 ================
var boo bool = true // false

// ============= 数值类型 ==============
// 整数
var a int = -1
var b uint = 1
var c int8 = 1

// 不同不可以互相复制
// c = a

// 浮点数
var f float64 = 1.1

// 指数
var d complex64 = 1 + 2i

// ============== 字符串 ==============
var str1 string = "str1"
var str2 string = `str2`
var str3 = str1 + str2

//  ============= 错误类型 ==============
var e error = errors.New("dem")

// 字符串是不可变的
// str1[0] = 'c' => err

func main() {
	// 修改字符串
	str := "asdasd"
	c := []byte(str)
	c[0] = 'c'
	newStr := string(c)
	fmt.Println(newStr)
}
