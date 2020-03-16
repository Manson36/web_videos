package api

import (
	"encoding/json"
	"github.com/web_videos/api/defs"
	"io"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter, errResp defs.ErrorResponse) {
	w.WriteHeader(errResp.HttpSC)

	resStr, _ := json.Marshal(errResp.Error)
	io.WriteString(w, string(resStr))
}

//这里为什么要将sc(StatusCode)单独洗出来？是因为这里的sc会随着业务的变化而变化，sc无法和error绑定
func sendNormalResponse(w http.ResponseWriter, resp string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}
