package dtos

type UserDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
