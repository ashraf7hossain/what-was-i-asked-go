package user 


type InputRegisterUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type InputLoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}