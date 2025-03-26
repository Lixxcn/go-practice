package main

import (
	"fmt"
)

// 使用comparable约束，允许比较操作
func Same[T comparable](a, b T) bool {
	return a == b
}

// 使用any约束，处理任意类型
func PrintType[T any](v T) {
	fmt.Printf("Type: %T, Value: %v\n", v, v)
}

func main() {
	// 使用comparable
	fmt.Println("4 = 3 is", Same(4, 3))           // 输出: false
	fmt.Println("aa = aa is", Same("aa", "aa"))   // 输出: true
	fmt.Println("4.1 = 4.15 is", Same(4.1, 4.15)) // 输出: false

	// 以下代码会编译错误，因为[]int不是comparable类型
	// _ = Same([]int{1, 2}, []int{1, 3})

	// 使用any
	PrintType(42)             // 输出: Type: int, Value: 42
	PrintType("hello")        // 输出: Type: string, Value: hello
	PrintType(3.14)           // 输出: Type: float64, Value: 3.14
	PrintType([]int{1, 2, 3}) // 输出: Type: []int, Value: [1 2 3]
}
