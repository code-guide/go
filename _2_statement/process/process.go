package main

import (
	"fmt"
)

func fun() {
here:
	fmt.Println("demo")

	goto here

}
func main() {
	// if
	a := 1
	if a > 0 {
		// 存在块级作用域
		name := "right"
		fmt.Println(name)
	} else if a == 0 {
		name := "0"
		fmt.Println(name)
	} else {
		name := "left"
		fmt.Println(name)
	}

	// ggto
	// fun()

	// for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	for i, j := 1, 2; i < 10; i++ {
		fmt.Println(i, j)
	}
	i := 0
	for i < 10 {
		i++
		fmt.Println(i)
	}

	// break continue
	for i := 0; i < 10; i++ {
		i++
		if i == 8 {
			break
		}
		if i == 3 {
			continue
		}
		fmt.Println("break & continue: ", i)
	}

	// for range
	demo := map[string]int{
		"one": 1,
		"two": 2,
	}

	for key, value := range demo {
		fmt.Println(key, ":", value)
	}

	// switch
	num := 1
	switch num {
	case 1, 2:
		fmt.Println(1, 2)
	case 3:
		fmt.Println(3)
		fallthrough
	case 4:
		fmt.Println(4)
	default:
		fmt.Println("demo")
	}
}
