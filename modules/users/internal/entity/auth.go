package entity

type AuthRegisterBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthLoginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthRegisterResponse struct {
	Token string `json:"token"`
}

type AuthLoginResponse struct {
	Token string `json:"token"`
}
