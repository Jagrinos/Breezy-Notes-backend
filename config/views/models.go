package views

type User struct {
	Id       string `json:"id"`
	Login    string `json:"login"`
	Email    string `json:"email"`
	About    string `json:"about"`
	Password string `json:"password"`
}

type UserNoId struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
	About    string `json:"about"`
}
