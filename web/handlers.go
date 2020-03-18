package web

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type HomePage struct {
	Name string
}

type UserPage struct {
	Name string
}

func apiHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method != http.MethodPost {
		re, _ := json.Marshal(ErrorRequestNotRecognized)
		io.WriteString(w, string(re))
		return
	}
	res, _ := ioutil.ReadAll(r.Body)
	apiBody := &ApiBody{}
	if err := json.Unmarshal(res, apiBody); err != nil {
		re, _ := json.Marshal(ErrorRequestBodyParseFailed)
		io.WriteString(w, string(re))
	}

	//前面是api处理的预处理，现在我们要真正开始处理request
}

func homeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//用户验证：通过cookie中的session验证
	cname, err1 := r.Cookie("username")
	sid, err2 := r.Cookie("session")
	if err1 != nil || err2 != nil {
		//进入home页面
		p := &HomePage{Name: "awen"}
		t, e := template.ParseFiles(".template/home.html")
		if e != nil {
			log.Println("Parsing template home.html error", e.Error())
			return
		}
		t.Execute(w, p)
	}
	if len(cname.Value) != 0 && len(sid.Value) != 0 {
		http.Redirect(w, r, "/userhome", http.StatusFound)
		return
	}
}

func userhomeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cname, err1 := r.Cookie("username")
	_, err2 := r.Cookie("session")
	if err1 != nil || err2 != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	fname := r.FormValue("username")
	var p *UserPage
	if len(cname.Value) != 0 {
		p = &UserPage{Name: cname.Value}
	} else if len(fname) != 0 { //如果在cookie中没有找到，那么我们在表单提交中去找
		p = &UserPage{Name: fname}
	}

	t, e := template.ParseFiles("./templates/userhome.html")
	if e != nil {
		log.Println("Parsing userhome.html error:", e.Error())
		return
	}

	t.Execute(w, p)
}

//可以建一个build.sh 文件
//作用是：web server在编译的时候，我们做两件事情，第一个用go先把它编译过来，第二个把它挪到一个真正的web ui的文件夹下
