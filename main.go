package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)
	return router
}

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "Create User Handler")
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}

func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":8080", r)

}

/*
listenANDServer流程：
listen -> RegisterHandlers -> handlers(CreateUser,Login)自动以goroutine 的方式启动
handler -> validation(校验){1.request是不是合法，2.user是不是合法用户} -> business logic(逻辑处理) -> response
validation：
{1.data model(数据结构) 2.error handling}
*/
