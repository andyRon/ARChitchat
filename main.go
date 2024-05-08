package main

import (
	. "github.com/andyron/architchat/routes"
	"log"
	"net/http"
)

func main() {
	startWebServer("8080")
}

func startWebServer(port string) {
	r := NewRouter()

	// 处理静态资源文件
	assets := http.FileServer(http.Dir("public"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", assets))

	http.Handle("/", r)

	log.Println("Starting HTTP service at " + port)
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}
