package Views

type User struct {
	Id       string `bson:"_id"`
	Login    string `bson:"login"`
	Password string `bson:"password"`
}

type UserWithoutId struct {
	Login    string `bson:"login"`
	Password string `bson:"password"`
}

type UserWithoutIdJson struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
