package main

import (
	"chitchat/data"
	"html/template"
	"net/http"
)

func main() {

	// net/http 标准库提供了一个默认的多路复用器，通过NewServeMux()函数创建
	// 创建一个多路复用器（多路转接器）
	mux := http.NewServeMux()
	// fmt.Println("mux=", mux) // mux= &{{{0 0} 0 0 0 0} map[] [] false}
	// fmt.Printf("%T", mux)    // *http.ServeMux

	// 使用FileServer函数创建一个能够为指定目录中的静态文件服务的处理器
	files := http.FileServer(http.Dir("/public"))
	// 并将这个处理器传递给多路复用器的Handel函数
	mux.Handle("/static/", http.StripPrefix("/static/", files)) // 使用StripPrefix函数移除请求URL中的指定前缀
	// 服务器接收到一个以/static/开头的URL请求时，以上两行代码会移除URL中的/static/
	// 然后在public目录中查找被请求的文件。

	// 为了将发送至根URL的请求重定向到处理器，使用HandleFunc函数
	// HandleFunc接受一个URL和一个处理器的名字作为参数
	mux.HandleFunc("/", index) // HandleFunc函数把请求重定向到处理器函数
	// server := &http.Server{
	// 	Addr:    "0.0.0.0:8080",
	// 	Handler: mux,
	// }
	// server.ListenAndServe()
}

// index是处理器函数
// 处理器函数实际上就是一个接受ResponseWriter和Request指针作为参数的Go函数
func index(w http.ResponseWriter, r *http.Request) {
	files := []string{"templates/layout.html",
		"templates/navbar.html",
		"templates/index.html"}
	templates := template.Must(template.ParseFiles(files...))
	threads, err := data.Threads()
	if err == nil {
		templates.ExecuteTemplate(w, "layout", threads)
	}
}
