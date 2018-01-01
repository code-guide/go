package main

import "fmt"

func SliceDemo() {
	slc := make([]int, 1, 4)

	for i := 0; i < 4; i++ {
		slc = append(slc, i)
	}

	fmt.Println(slc)
}
