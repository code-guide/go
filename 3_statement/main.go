package main

import "fmt"

func main() {
	// if 语句
	if 2 > 1 {
		fmt.Println("2 > 1")
	}

	// switch 语句
	score := 80
	switch {
	case score >= 90:
		fmt.Println("优秀")
	case score >= 80:
		fmt.Println("良")
	case score >= 60:
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}

	// for 循环
	fmt.Println("for 循环：")
	for i := 4; i < 10; i++ {
		fmt.Println(i)
	}

	// 使用for实现while语句
	fmt.Println("for 循环实现while语句：")
	i := 0
	for i < 10 {
		fmt.Println(i)
		i = i + 2
	}

	// 无限循环
	fmt.Println("for无限循环：")
	j := 0
	for {
		fmt.Println(j)
		j = j + 2

		if j > 10 {
			break
		}
	}

	// for range遍历
	for item, index := range []int{1, 3, 5} {
		fmt.Println(item, index)
	}
}
