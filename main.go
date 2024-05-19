package main

import (
	. "github.com/andyron/architchat/config"
	. "github.com/andyron/architchat/routes"
	"log"
	"net/http"
)

func main() {
	startWebServer()
}

// 通过指定端口启动 Web 服务器
func startWebServer() {
	// 在入口位置初始化全局配置
	config := LoadConfig()

	r := NewRouter()

	// 处理静态资源文件
	// 初始化文件服务器和目录为当前目录下的 public 目录
	assets := http.FileServer(http.Dir(config.App.Static))
	// 指定静态资源路由及处理逻辑：将 /static/ 前缀的 URL 请求去除 static 前缀，然后在文件服务器查找指定文件路径是否存在（public 目录下的相对地址）。
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", assets))

	http.Handle("/", r)

	log.Println("Starting HTTP service at " + config.App.Address)
	err := http.ListenAndServe(config.App.Address, nil)

	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + config.App.Address)
		log.Println("Error: " + err.Error())
	}
}
