package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/web_videos/api/dbops"
	"github.com/web_videos/api/defs"
	"github.com/web_videos/api/session"
	"io"
	"io/ioutil"
	"net/http"
)

/*
Handler任务：
1.从request中获取body，并将其反序列化到创建好的 Handler需要的ReqBody结构
2.对传入的参数进行验证
3.执行数据库操作和session操作
4.将执行的结果返回到response中
*/
func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//io.WriteString(w, "Create User Handler")
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}

	if err := json.Unmarshal(res, &ubody); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	if err := dbops.AddUserCredential(ubody.UserName, ubody.Pwd); err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
	}

	id := session.GenerateNewSessionID(ubody.UserName)
	su := defs.SignedUp{Success: true, SessionId: id}

	if resp, err := json.Marshal(su); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp), 201)
	}
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}
