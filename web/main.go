package web

import (
	"github.com/julienschmidt/httprouter"
	"github.com/web_videos/api"
	"net/http"
)

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()

	//为什么这里一个path使用了两个方法：登录页面logout之后会有页面跳转
	router.GET("/", homeHandler) //主页
	router.POST("/", homeHandler)

	router.GET("/userhome", userhomeHandler)
	router.POST("/userhome", userhomeHandler)

	router.POST("/api", apiHandler())
	//file server,放静态文件的server。在goland中它提供原生的能够把一个url绑到一个文件夹下，作为他的一个file server
	//效果就是：127.0.0.1:8080/static 直接跳转到template文件夹下，可以访问它的子文件
	router.ServeFiles("static/*filepath", http.Dir("./template"))

	return router
}

func main() {
	r := RegisterHandler()
	http.ListenAndServe(":8080", r)
}
