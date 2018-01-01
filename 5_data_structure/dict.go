package main

import "fmt"

func DictDemo() {
	dict := make(map[string]int) // 声明
	dict["tom"] = 10             // 赋值
	TomAge, ok := dict["tom"]    // 取值，ok-idiom
	delete(dict, "tom")          // 删除

	fmt.Println("Tom:", TomAge, ok)
}
