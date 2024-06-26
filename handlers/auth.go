package handlers

import (
	"fmt"
	"github.com/andyron/architchat/models"
	"net/http"
)

// 用户认证相关处理器

// Login GET /login
func Login(w http.ResponseWriter, r *http.Request) {
	//t := parseTemplateFiles("auth.layout", "navbar", "login")
	//t.Execute(w, nil)
	generateHTML(w, nil, "auth.layout", "navbar", "login")
}

// Signup 注册页面 GET /signup
func Signup(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "auth.layout", "navbar", "signup")
}

// SingupAccount 注册新用户 POST /signup
func SingupAccount(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		//fmt.Println("Cannot parse form")
		danger(err, "Cannot parse form")
	}
	user := models.User{
		Name:     r.PostFormValue("name"),
		Email:    r.PostFormValue("email"),
		Password: r.PostFormValue("password"),
	}
	if err := user.Create(); err != nil {
		//fmt.Println("Cannot create user")
		danger(err, "Cannot create user")
	}
	http.Redirect(w, r, "/login", 302)
}

// Authenticate POST /authenticate 通过邮箱和密码字段对用户进行认证
func Authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	user, err := models.UserByEmail(r.PostFormValue("email"))
	if err != nil {
		//fmt.Println("Cannot find user")
		danger(err, "Cannot find user")
	}
	if user.Password == models.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			//fmt.Println("Cannot create session")
			danger(err, "Cannot create session")
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

// Logout GET /logout
func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_cookie")
	if err != http.ErrNoCookie {

		fmt.Println("Failed to get cookie", cookie.Value)

		warning(err, "Failed to get cookie")
		session := models.Session{Uuid: cookie.Value}

		session.DeleteByUUID()
	}
	http.Redirect(w, r, "/", 302)
}
