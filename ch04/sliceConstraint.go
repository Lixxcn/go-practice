package main

import (
	"fmt"
)

// f1 使用完整的interface{}语法定义约束
// S 的底层类型必须是切片类型，且元素类型为E
// E 可以是任意类型
// ~[]E 表示任何底层类型为[]E的类型，包括[]E本身和基于[]E定义的类型别名
func f1[S interface{ ~[]E }, E interface{}](x S) int {
	return len(x)
}

// f2 使用简化语法定义约束
// S 的底层类型必须是切片类型，且元素类型为E
// E 可以是任意类型
// ~[]E 表示任何底层类型为[]E的类型
// 与f1功能相同，但语法更简洁
func f2[S ~[]E, E interface{}](x S) int {
	return len(x)
}

// f3 使用any代替interface{}
// S 的底层类型必须是切片类型，且元素类型为E
// E 可以是任意类型
// ~[]E 表示任何底层类型为[]E的类型
// 与f1、f2功能相同，但使用更现代的语法
func f3[S ~[]E, E any](x S) int {
	return len(x)
}

func main() {
	// 定义自定义切片类型
	type MyIntSlice []int
	type MyFloatSlice []float64

	// 调用f1，传入标准[]int
	fmt.Println("Len:", f1([]int{1, 2, 3})) // 输出: 3
	// 调用f1，传入自定义类型MyIntSlice
	fmt.Println("Len:", f1(MyIntSlice{1, 2, 3})) // 输出: 3

	// 调用f2，传入标准[]float64
	fmt.Println("Len:", f2([]float64{1.1, -2})) // 输出: 2
	// 调用f2，传入自定义类型MyFloatSlice
	fmt.Println("Len:", f2(MyFloatSlice{1.1, -2})) // 输出: 2

	// 调用f3，传入标准[]float32
	fmt.Println("Len:", f3([]float32{1.1, -2})) // 输出: 2
}
