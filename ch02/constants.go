package main

import (
	"fmt"
)

type Digit int
type Power2 int

const PI = 3.1415926

const (
	C1 = "C1C1C1"
	C2 = "C2C2C2"
	C3 = "C3C3C3"
)

func main() {
	const s1 = 123
	var v1 float32 = s1 * 12
	fmt.Println(v1)
	fmt.Println(PI)

	const (
		Zero Digit = iota
		One
		Two
		Three
		Four
	)

	fmt.Println(One)
	fmt.Println(Two)

	const (
		p2_0 Power2 = 1 << iota // 1 << 0 = 1
		_                       // 1 << 1 = 2 (被忽略)
		p2_2                    // 1 << 2 = 4
		_                       // 1 << 3 = 8 (被忽略)
		p2_4                    // 1 << 4 = 16
		_                       // 1 << 5 = 32 (被忽略)
		p2_6                    // 1 << 6 = 64
	)

	fmt.Println("2^0:", p2_0)
	fmt.Println("2^2:", p2_2)
	fmt.Println("2^4:", p2_4)
	fmt.Println("2^6:", p2_6)
}
