package input

type MessageInput struct {
	Sender   string `json:"sender" validate:"required,notblank" xml:"sender" form:"sender" query:"sender"`
	Receiver string `json:"receiver" validate:"required,notblank" xml:"receiver" form:"receiver" query:"receiver"`
	Message  string `json:"message" xml:"message" form:"message" query:"message"`
}
