package routes

import (
	"github.com/andyron/architchat/handlers"
	"net/http"
)

// WebRoute 具体路由
type WebRoute struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type WebRoutes []WebRoute

// 定义所有Web路由
var webRoutes = WebRoutes{
	{
		"home",
		"GET",
		"/",
		handlers.Index,
	},
	{
		"signup",
		"GET",
		"/signup",
		handlers.Signup,
	},
	{
		"signupAccount",
		"POST",
		"/signup_account",
		handlers.SingupAccount,
	},
	{
		"login",
		"GET",
		"/login",
		handlers.Login,
	},
	{
		"auth",
		"POST",
		"/authenticate",
		handlers.Authenticate,
	},
	{
		"logout",
		"GET",
		"/logout",
		handlers.Logout,
	},

	{
		"newThread",
		"GET",
		"/thread/new",
		handlers.NewThread,
	},
	{
		"createThread",
		"POST",
		"/thread/create",
		handlers.CreateThread,
	},
	{
		"readThread",
		"GET",
		"/thread/read",
		handlers.ReadThread,
	},

	{
		"postThread",
		"POST",
		"/thread/post",
		handlers.PostThread,
	},

	{
		"error",
		"GET",
		"/err",
		handlers.Err,
	},
}
