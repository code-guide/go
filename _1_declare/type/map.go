/**
 * 引用类型 - 字典 map
 */
package main

import (
	"fmt"
)

func main() {
	// 声明方式 1
	var numbers1 map[string]int

	// 声明方式 2
	numbers2 := make(map[string]int)

	// 声明方式 3
	number3 := map[string]int{
		"one": 1,
	}

	// 赋值
	numbers2["one"] = 1

	fmt.Println(numbers1, numbers2, number3)
}
