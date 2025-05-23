package utils

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func Success(result interface{}) *Response {
	return &Response{
		Code:    0,
		Message: "success",
		Result:  result,
	}
}

func Error(code int, msg string, err error) *Response {
	return &Response{
		Code:    code,
		Message: msg,
		Result:  err,
	}
}
