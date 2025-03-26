package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	const sL = "\x99\x00ab\x50\x00\x23\x50\x29\x9c李"
	fmt.Println(len(sL))
	fmt.Println(utf8.RuneCountInString(sL))

	for i := 0; i < len(sL); i++ {
		if unicode.IsPrint(rune(sL[i])) {
			fmt.Printf("%c\n", sL[i])
		} else {
			fmt.Println("Not printable!")
		}
	}

	fmt.Println("Printing as runes")

	for _, r := range sL {
		if unicode.IsPrint(r) {
			fmt.Printf("%c\n", r)
		} else {
			fmt.Println("Not printable!")
		}
	}
}
