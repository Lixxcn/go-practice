package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 空map
	emptyMap := map[string]int{}
	// nil map
	var nilMap map[string]int

	fmt.Println("emptyMap:", emptyMap) // 输出: map[]
	fmt.Println("nilMap:", nilMap)     // 输出: map[]

	// 使用反射检查
	fmt.Println("emptyMap is nil:", reflect.ValueOf(emptyMap).IsNil()) // 输出: false
	fmt.Println("nilMap is nil:", reflect.ValueOf(nilMap).IsNil())     // 输出: true

	// 安全操作
	val, exists := nilMap["test"]
	fmt.Println("nilMap lookup:", val, exists) // 输出: 0 false

	// 危险操作（会导致panic）
	// nilMap["test"] = 1
}
