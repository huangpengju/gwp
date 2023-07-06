// main包，Go语言要求可执行程序必须位于main包当中
package main

import (
	"fmt"
	"net/http"
)

// handler 处理器函数接受2个参数
// http.ResponseWriter
// ResponseWriter接口被HTTP处理器用于构造HTTP回复。(回复的内容是“Hello World, %s!”)
// func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
// Fprintf根据format参数生成格式化的字符串并写入w。返回写入的字节数和遇到的任何错误。
// 在这里使用 Fprintf 函数对I/O格式化,把“Hello World, %s”写入ResponseWriter接口
func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
	// fmt.Printf("a= %T \n", request.URL.Path)
	// fmt.Println("b=", request.URL.Path)
	// fmt.Println("b=", request.URL.Path[1:])
}

// Go语言规定，每个需要被编译为二进制可执行文件的程序都必须包含main函数
// 用作程序执行时的起点：
func main() {
	// HandleFunc注册一个处理器函数handler和对应的模式pattern（注册到DefaultServeMux）。ServeMux的文档解释了模式的匹配机制。
	// 首先把之前定义的 handler 函数设置成根(root) URL （／）被访问时的处理器，然
	http.HandleFunc("/", handler)
	// 然后启动服务器并让它监昕系统的 8080 端口
	http.ListenAndServe(":8080", nil)
}
