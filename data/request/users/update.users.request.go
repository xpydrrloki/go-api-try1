package request

type UpdateUserRequest struct {
	Id       int    `validate:"required,min=1,max=200" json:"id"`
	Username string `validate:"required,min=1,max=200" json:"username"`
	Email    string `validate:"required,min=1,max=200" json:"email"`
	Password string `validate:"required,min=6,max=200" json:"password"`
}
