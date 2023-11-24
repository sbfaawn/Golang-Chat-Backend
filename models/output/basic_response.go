package output

type BaseResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
	Error   string `json:"error"`
}
