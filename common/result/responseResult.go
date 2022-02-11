package result

type ResponseSuccessModel struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type RepsonseErrorModel struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
}

type NullJson struct {
}

//正常结果
func Success(data interface{}) *ResponseSuccessModel {
	return &ResponseSuccessModel{200, "OK", data}
}

//错误结果
func Error(code uint32, msg string) *RepsonseErrorModel {
	return &RepsonseErrorModel{code, msg}
}
