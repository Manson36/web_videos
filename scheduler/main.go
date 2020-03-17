package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/web_videos/scheduler/task_runner"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/video-delte-record/:vid-id", videoDelRecRecord)
	return router
}

func main() {
	//在这里http.ListenAndServer本身就是阻塞的，如果将那行代码注释掉，要执行下面的go ...需要添加阻塞：
	//for{} 或c := make(chan int);go ... ; <-c
	go task_runner.Start()
	r := RegisterHandlers()
	http.ListenAndServe(":9001", r)
}
