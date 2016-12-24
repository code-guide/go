package main

// type 定义函数类型，相同类型：同名同参同返回值
type demo func(int) int

func main() {
	say(dd)
}

func dd(num int) int {
	return num
}

func say(cb demo) {
	num := 10
	cb(num)
}
