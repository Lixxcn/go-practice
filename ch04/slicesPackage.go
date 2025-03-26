package main

import (
	"fmt"
	"slices"
)

func main() {
	// 创建一个初始切片
	s1 := []int{1, 2, -1, -2}

	// Clone: 创建切片的完整副本
	s2 := slices.Clone(s1)
	// Clone: 创建切片部分范围的副本
	s3 := slices.Clone(s1[2:])
	fmt.Println(s1[2], s2[2], s3[0]) // 输出: -1 -1 -1

	// 修改原始切片
	s1[2] = 0
	s1[3] = 0
	fmt.Println(s1[2], s2[2], s3[0]) // 输出: 0 -1 -1

	// Compact: 去除连续重复的元素
	s1 = slices.Compact(s1)
	fmt.Println("s1 (compact):", s1) // 输出: [1 2 0]

	// Contains: 检查切片是否包含特定元素
	fmt.Println(slices.Contains(s1, 2), slices.Contains(s1, -2)) // 输出: true false

	// 创建一个长度为10，容量为100的切片
	s4 := make([]int, 10, 100)
	fmt.Println("Len:", len(s4), "Cap:", cap(s4)) // 输出: Len: 10 Cap: 100

	// Clip: 将切片的容量缩减到与长度相同
	s4 = slices.Clip(s4)
	fmt.Println("Len:", len(s4), "Cap:", cap(s4)) // 输出: Len: 10 Cap: 10

	// Min/Max: 查找切片中的最小值和最大值
	fmt.Println("Min:", slices.Min(s1), "Max:", slices.Max(s1)) // 输出: Min: 0 Max: 2

	// Replace: 替换切片中的元素
	// 将s2[1]和s2[2]替换为100和200
	s2 = slices.Replace(s2, 1, 3, 100, 200)
	fmt.Println("s2 (replaced):", s2) // 输出: [1 100 200 -2]

	// Sort: 对切片进行排序
	slices.Sort(s2)
	fmt.Println("s2 (sorted):", s2) // 输出: [-2 1 100 200]
}
