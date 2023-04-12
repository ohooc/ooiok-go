package main

import (
	"log"
	"os"
)

func main() {

	// 设置前缀
	log.SetPrefix("prefix_")

	// 设置打印样式
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)

	file, err := os.OpenFile("./go-logger/go-log/base/base.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		return
	}
	log.SetOutput(file)

	// Print 系列都会换行
	log.Print("go hello...")
	log.Printf("hello %s", "golang")
	log.Println("hello go...")

	log.Fatal("打印并退出")

	log.Panic("打印日志，及其详细信息，然后退出")
}
