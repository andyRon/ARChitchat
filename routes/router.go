package routes

import "github.com/gorilla/mux"

// NewRouter 返回一个mux.Router类型指针，从而可以当作路由处理器使用
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range webRoutes {
		// 将每个web路由应用到路由器
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}
