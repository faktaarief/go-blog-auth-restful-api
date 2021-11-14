package helper

type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password"`
}
