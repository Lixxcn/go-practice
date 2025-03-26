package main

import (
	"log"
	"log/syslog"
)

func main() {
	// 创建一个新的syslog连接
	// LOG_SYSLOG表示使用系统日志设施
	// "systemLog.go"是日志消息的前缀
	sysLog, err := syslog.New(syslog.LOG_SYSLOG, "systemLog.go")

	// 检查是否成功连接到syslog
	if err != nil {
		log.Println(err) // 如果连接失败，打印错误信息
		return
	} else {
		// 将标准日志输出重定向到syslog
		log.SetOutput(sysLog)
		// 发送一条日志消息到syslog
		log.Print("Everything is fine!")
	}
}
