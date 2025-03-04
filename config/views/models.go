package views

type User struct {
	Id       string `json:"id"`
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
	About    string `json:"about"`
}

type UserNoId struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
	About    string `json:"about"`
}
