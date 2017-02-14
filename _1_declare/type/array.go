/**
 * 引用类型 - 数组 array
 */
package main

import (
	"fmt"
)

func change(arr [3]int) {
	arr[0] = 3
}

func main() {
	// =========== 数组 array =============
	// 声明1 : var arrName [n]type
	// 声明时不可赋值
	var arr [3]int
	arr[0] = 1
	arr[0] = 2
	arr[1] = 2

	// 赋值不可越界
	// arr[4] = 1

	// 长度也是类型
	// var arr2 [4]int
	// arr2 = arr

	// 数组之间赋值是值传递
	change(arr)
	fmt.Println(arr)

	// 声明2 ： arrName := [n]type{items}
	arr2 := [3]int{1, 2, 3}
	arr3 := [10]int{1, 2, 3}
	arr4 := [...]int{1, 2, 3}
	fmt.Println(arr2, arr3, arr4)

	// 支持多维数组
	arr5 := [2][3]int{{1, 2, 3}, {2, 3, 4}}
	fmt.Println(arr5)
}
