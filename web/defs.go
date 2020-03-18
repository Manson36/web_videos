package web

//我们可以看到，在之前的api可以分为Get、Post、delete几个部分，现在将他们统一为一种方式
type ApiBody struct {
	Url     string `json:"url"`
	Method  string `json:"method"`
	ReqBody string `json:"reqBody"`
}

type Err struct {
	Error   string `json:"error"`
	ErrCode string `json:"err_code"`
}

var (
	ErrorRequestNotRecognized   = Err{Error: "api not recognized, bad request", ErrCode: "001"}
	ErrorRequestBodyParseFailed = Err{Error: "request body is not correct", ErrCode: "002"}
	ErrorInternalFaults         = Err{Error: "internal error", ErrCode: "003"}
)
