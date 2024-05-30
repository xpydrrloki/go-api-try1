package request

type CreateUsersRequest struct {
	Username string `validate:"required,min=1,max=200" json:"username"`
	Email    string `validate:"required,min=1,max=200" json:"email"`
	Password string `validate:"required,min=6,max=200" json:"password"`
}
