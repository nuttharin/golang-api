package request

type UserReq struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserUpdateReq struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
