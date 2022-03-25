package param

type Users struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserModify struct {
	NewPassword string `json:"new_password"`
}
