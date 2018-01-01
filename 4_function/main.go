package main

import (
	"errors"
	"fmt"
)

func add(x, y int) int {
	return x + y
}

// 多返回值
func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("division by zero")
	}

	return x / y, nil
}

// 高阶函数
func run(cb func()) {
	cb()
}

func main() {
	// 延时调用：无论是否发生错误，都最后调用，采用栈的方式
	defer func() {
		fmt.Println("end")
	}()

	result, err := divide(2, 0)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("2 / 1 =", result)

	run(func() {
		fmt.Println("cb running...")
	})
}
