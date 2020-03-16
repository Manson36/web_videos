package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

/*
我们先说一下streaming的过程应该是什么样的，在我们这种程度上还有好多方法来streaming整个过程：
1.我们自己将videos的内容格式化成它的二进制的data stream，用stream的方式给它传到client端，这样二进制的速度和带宽我们都可以控制，但是实现起来非常复杂
2.今天我们使用简单的方法，也是在web server中通用的方法：
*/
func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	vl := VIDEO_DIR + vid //video link

	video, err := os.Open(vl)
	if err != nil {
		log.Println("Error when try to open file", err.Error())
		sendErrorResponse(w, http.StatusInternalServerError, "Internal error")
	}

	//这里的Header.Set()是 我们需要把的返回的response的请求给他加一个Header的强制提醒，叫Content-Type。
	//因为我们在传这个文件的时候，这个文件可能是没有它的扩展名(extention)的，然而它的文件里面真正的二进制码一样是视频的MP4格式，
	//这时在client端会自动将它作为video/mp4这个格式来解析，解析之后，就会将他组成真正的视频来播放
	w.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(w, r, "", time.Now(), video)
	defer video.Close()
}

func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//这个函数是设定io最大能读取的文件的大小，单位是byte(如果是bit还要乘 8)
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "File is too big.")
		return
	}

	file, _, err := r.FormFile("file") //<form name = "file">
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Internal error")
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("read file error ", err.Error())
		sendErrorResponse(w, http.StatusInternalServerError, "Internal ")
	}

	fileName := p.ByName("vid-id")
	err = ioutil.WriteFile(VIDEO_DIR+fileName, data, 0666)
	if err != nil {
		log.Println("Write file error:", err.Error())
		sendErrorResponse(w, http.StatusInternalServerError, "internal error")
		return
	}

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "upload successfully")
}

//之后会细讲的template的东西，go做web页面非常有用的东西 "html/template"
func testPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, _ := template.ParseFiles("./videos/upload.html")
	t.Execute(w, nil)
}
