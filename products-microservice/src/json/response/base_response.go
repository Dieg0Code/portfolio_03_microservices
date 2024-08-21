package response

type BaseResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}
