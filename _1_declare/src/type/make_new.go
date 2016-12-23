package main

import (
	"fmt"
)

func main() {
	// new  => 返回内存指针
	var pointer *int = new(int)
	*pointer = 1

	// make => 返回初始化后的值
	var mapData = make(map[string]int)

	fmt.Println(pointer, *pointer, mapData)
}
