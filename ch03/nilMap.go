package main

import (
	"fmt"
)

func main() {
	// 创建一个空的map
	aMap := map[string]int{}
	fmt.Println("aMap:", aMap) // 输出: map[]
	aMap["test"] = 1
	fmt.Println("aMap:", aMap) // 输出: map[test:1]

	// 将map设置为nil
	aMap = nil
	fmt.Println("aMap:", aMap) // 输出: map[]

	// 检查map是否为nil
	if aMap == nil {
		fmt.Println("nil map!") // 输出: nil map!
		// 重新初始化map
		aMap = map[string]int{}
	}
	fmt.Println("aMap:", aMap) // 输出: map[]
	aMap["test"] = 1
	fmt.Println("aMap:", aMap) // 输出: map[test:1]

	// 这将导致程序崩溃！
	aMap = nil
	aMap["test"] = 1
}
