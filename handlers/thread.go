package handlers

import (
	"github.com/andyron/architchat/models"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"net/http"
)

// 群组

// NewThread GET /threads/new 创建群组页面
func NewThread(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		generateHTML(w, nil, "layout", "auth.navbar", "new.thread")
	}
}

// CreateThread POST /thread/create 创建群组
func CreateThread(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err = r.ParseForm()
		if err != nil {
			//fmt.Println("Cannot parse form")
			danger(err, "Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			//fmt.Println("Cannot get user from session")
			danger(err, "Cannot get user from session")
		}
		topic := r.PostFormValue("topic")
		if _, err := user.CreateThread(topic); err != nil {
			//fmt.Println("Cannot create thread")
			danger(err, "Cannot create thread")
		}
		http.Redirect(w, r, "/", 302)
	}
}

// ReadThread GET /thread/read 通过ID渲染指定群组页面
func ReadThread(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	uuid := vals.Get("id")
	thread, err := models.ThreadByUUID(uuid)
	if err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			MessageID: "thread_not_found",
		})
		errorMessage(w, r, msg)
	} else {
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, &thread, "layout", "navbar", "thread")
		} else {
			generateHTML(w, &thread, "layout", "auth.navbar", "auth.thread")
		}
	}
}
