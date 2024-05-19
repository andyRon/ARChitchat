package handlers

import (
	"github.com/andyron/architchat/models"
	"net/http"
)

// Index 论坛首页路由处理器方法
func Index(w http.ResponseWriter, r *http.Request) {
	threads, err := models.Threads()
	if err == nil {
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, threads, "layout", "navbar", "index")
		} else {
			generateHTML(w, threads, "layout", "auth.navbar", "index")
		}
	}
}

// Err 全局的、渲染错误页面的处理器
func Err(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, vals.Get("msg"), "layout", "navbar", "error")
	} else {
		generateHTML(w, vals.Get("msg"), "layout", "auth.navbar", "error")
	}
}
