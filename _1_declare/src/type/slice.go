/**
 * 引用类型 - 切片 slice
 */
package main

import (
	"fmt"
)

func change(s []int) {
	s[1] = 10
}

func main() {
	// 声明方式  与array相同，不需要长度
	var s1 []int
	s2 := []int{1, 2, 3}

	fmt.Println(s1, s2)

	a := s2[1:]
	b := s2[:]
	fmt.Println(a, b)

	// 引用类型
	change(a)
	fmt.Println(s2)
}
