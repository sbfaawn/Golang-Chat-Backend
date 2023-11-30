package input

type AccountInput struct {
	Email    string `json:"email" validate:"notblank,email" xml:"email" form:"email" query:"email"`
	Username string `json:"username" validate:"required,notblank" xml:"username" form:"username" query:"username"`
	Password string `json:"password" validate:"required,notblank,password" xml:"password" form:"password" query:"password"`
}
