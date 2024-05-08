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
	assets := http.FileServer(http.Dir(config.App.Static))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", assets))

	http.Handle("/", r)

	log.Println("Starting HTTP service at " + config.App.Address)
	err := http.ListenAndServe(config.App.Address, nil)

	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + config.App.Address)
		log.Println("Error: " + err.Error())
	}
}
