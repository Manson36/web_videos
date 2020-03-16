package api

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

/*首先演示一下添加MiddleWare部分：
1.在main中直接就是RegisterHandlers，它返回的是*httprouter.Router，实际上它是实现了http中的一个接口 http.Handler
2.http.Handler接口 只有一个方法ServeHTTP(http.ResponseWriter, *http.Request)，duck type，只要实现了ServeHTTP 即可实现http.Handler
*/

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//check session
	validateUser(w, r)

	m.r.ServeHTTP(w, r)
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)
	return router
}

func main() {
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	http.ListenAndServe(":8080", mh)

}

/*
listenANDServer流程：
listen -> RegisterHandlers -> handlers(CreateUser,Login)自动以goroutine 的方式启动
handler -> validation(校验){1.request是不是合法，2.user是不是合法用户} -> business logic(逻辑处理) -> response
validation：
{1.data model(数据结构) 2.error handling}
*/
