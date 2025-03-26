package main

import (
	"fmt"
	"os"
	"regexp"
)

func matchInt(s string) bool {
	t := []byte(s)
	fmt.Println("t:", t)
	re := regexp.MustCompile(`^[-+]?\d+$`)
	fmt.Println("re:", re)
	return re.Match(t)
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Usage: <utility> string.")
		return
	}

	s := arguments[1]
	ret := matchInt(s)
	fmt.Println(ret)
}
