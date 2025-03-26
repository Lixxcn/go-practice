package main

// https://stackoverflow.com/questions/25159236/panic-when-compiling-a-regular-expression

import (
	"fmt"
	"regexp"
)

func main() {
	// 这是一个原始字符串字面量，用于密码强度验证
	// 要求：至少7个字符，且包含至少一个数字
	var re string = `^.*(?=.{7,})(?=.*\d)$`

	// 使用regexp.Compile编译正则表达式
	// 适合场景：当正则表达式可能包含错误时（如用户输入）
	// 优点：可以优雅地处理错误
	exp1, err := regexp.Compile(re)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("RegExp:", exp1)

	// 使用regexp.MustCompile编译正则表达式
	// 适合场景：当正则表达式是硬编码的，且确定不会出错时
	// 优点：代码更简洁，但如果正则表达式错误会直接panic
	exp2 := regexp.MustCompile(re)
	fmt.Println(exp2)
}
