package handlers

import (
	"github.com/andyron/architchat/models"
	"net/http"
)

// 用户认证相关处理器

// GET /login
func Login(w http.ResponseWriter, r *http.Request) {
	//t := parseTemplateFiles("auth.layout", "navbar", "login")
	//t.Execute(w, nil)
	generateHTML(w, nil, "auth.layout", "navbar", "login")
}

// GET /signup
// 注册页面
func Signup(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, nil, "auth.layout", "navbar", "signup")
}

// POST /signup
// 注册新用户
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

// POST /authenticate
// 通过邮箱和密码字段对用户进行认证
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

// GET /logout
func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_cookie")
	if err != http.ErrNoCookie {
		//fmt.Println("Failed to get cookie")
		warning(err, "Failed to get cookie")
		session := models.Session{Uuid: cookie.Value}
		session.DeleteByUUID()
	}
	http.Redirect(w, r, "/", 302)
}
