/**
 * 函数 - defer
 */
package main

import (
	"fmt"
)

func main() {
	Demo()
}

func Demo() {
	for i := 0; i < 10; i++ {
		// 后入先出
		defer fmt.Println(i)
	}
}
