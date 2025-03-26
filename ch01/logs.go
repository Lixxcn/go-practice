package main

/*
log.Fatal():
打印日志信息后，调用os.Exit(1)立即终止程序。
不会执行defer语句。
通常用于遇到无法恢复的错误时，立即停止程序运行。
log.Panic():
打印日志信息后，调用panic()函数。
会执行defer语句。
通常用于遇到严重错误时，但希望程序能够执行清理操作后再崩溃。
*/
import (
	"log"
	"os"
)

func main() {
	if len(os.Args) != 1 {
		log.Fatal("Fatal: Hello World!")
	}
	log.Panic("Panic: Hello World!")
}
