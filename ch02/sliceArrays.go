package main

import (
	"fmt"
)

/*
**关键点解释**：
1. **切片与数组的关系**：
  - 切片是对数组的一个"视图"或"引用"
  - 切片本身不存储数据，它只是指向底层数组的一部分
  - 修改切片会修改底层数组

2. **切片的组成**：
  - 指针：指向底层数组的某个元素
  - 长度：切片中元素的数量
  - 容量：从切片开始位置到底层数组末尾的元素数量

3. **append操作的影响**：
  - 当append操作不超过切片的容量时，会修改底层数组
  - 当append操作超过切片的容量时，Go会创建一个新的底层数组，并将原数据复制过去
  - 此时，新的切片与原来的数组不再共享数据

4. **代码中的具体示例**：
  - 初始时，S0和S12都共享数组a
  - 当S0通过append操作超过容量后，它使用新的底层数组
  - 此时，修改数组a不会影响S0，但会影响S12
*/
func change(s []string) {
	s[0] = "Change_function"
}

func main() {
	// 创建一个包含4个元素的数组
	a := [4]string{"Zero", "One", "Two", "Three"}
	fmt.Println("a:", a)

	// 创建一个指向数组a第一个元素的切片
	var S0 = a[0:1]
	fmt.Println(S0)
	S0[0] = "S0" // 修改切片会修改底层数组

	// 创建一个指向数组a第二个和第三个元素的切片
	var S12 = a[1:3]
	fmt.Println(S12)
	S12[0] = "S12_0" // 修改切片会修改底层数组
	S12[1] = "S12_1" // 修改切片会修改底层数组

	fmt.Println("a:", a) // 可以看到数组a已经被修改

	// 通过函数修改切片
	change(S12)
	fmt.Println("a:", a) // 数组a再次被修改

	// 打印S0的容量和长度
	fmt.Println("Capacity of S0:", cap(S0), "Length of S0:", len(S0))
	fmt.Println("S0:", S0)

	// 向S0追加元素
	S0 = append(S0, "N1")
	S0 = append(S0, "N2")
	S0 = append(S0, "N3")
	a[0] = "-N1" // 修改数组a，这会影响S0
	fmt.Println("S0:", S0)
	fmt.Println("Capacity of S0:", cap(S0), "Length of S0:", len(S0))

	// 当追加元素超过容量时，Go会创建一个新的底层数组
	S0 = append(S0, "N4")

	fmt.Println("Capacity of S0:", cap(S0), "Length of S0:", len(S0))
	// 这个修改不会影响S0，因为S0现在使用新的底层数组
	a[0] = "-N1-"

	// 这个修改会影响S12，因为S12仍然使用原始数组a
	a[1] = "-N2-"

	fmt.Println("S0:", S0)
	fmt.Println("a: ", a)
	fmt.Println("S12:", S12)
}
