package web

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

//这个client.go就相当于于是一个代理的过程。
//再多说两句：我们在这边的大前端它等于是一个很薄的一层，这一层只做两件事情：1.转发业务的请求2.把它用来模板化的东西真正的通过网页渲染到前端
//转发业务请求有两种模式：proxy模式和api模式
//为什么有时候需要使用proxy透传：因为在很多的情况下我们使用api，通过api的包装的方法是是处理不了原生的http请求的，比如说在这里我们有原生的
//http请求(stream_server中的upload),里面我们需要传过来的是一个文件，我们在这里如果使用api透传的方法，先把要传的东西转换为api包装的形式，
//文件无法放到它的body中，这时候我们就需要直接代理的方法，Proxy模式

var httpClient *http.Client

func init() {
	httpClient = &http.Client{}
}

func request(b *ApiBody, w http.ResponseWriter, r *http.Request) {
	var resp *http.Response
	var err error

	switch b.Method {
	case http.MethodGet:
		req, _ := http.NewRequest("GET", b.Url, nil)
		req.Header = r.Header
		resp, err = httpClient.Do(req)
		if err != nil {
			log.Println(err)
			return
		}
		normalRequest(w, resp)

	case http.MethodPost:
		req, _ := http.NewRequest("POST", b.Url, bytes.NewBuffer([]byte(b.ReqBody)))
		req.Header = r.Header
		resp, err = httpClient.Do(req)
		if err != nil {
			log.Println(err)
			return
		}
		normalRequest(w, resp)
	case http.MethodDelete:
		req, _ := http.NewRequest("DELETE", b.Url, nil)
		req.Header = r.Header
		resp, err = httpClient.Do(req)
		if err != nil {
			log.Println(err)
			return
		}
		normalRequest(w, resp)
	default:
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "bad api request")
		return
	}

}

func normalRequest(w http.ResponseWriter, res *http.Response) {
	rec, err := ioutil.ReadAll(res.Body)
	if err != nil {
		re, _ := json.Marshal(ErrorInternalFaults)
		w.WriteHeader(500)
		io.WriteString(w, string(re))
		return
	}

	w.WriteHeader(res.StatusCode)
	io.WriteString(w, string(rec))
}
