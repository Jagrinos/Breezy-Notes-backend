package models

type User struct {
	Id       string `bson:"_id"`
	Login    string `bson:"login"`
	Password string `bson:"password"`
}

type UserWithoutId struct {
	Login    string `bson:"login" json:"login"`
	Password string `bson:"password" json:"password"`
}

type Block struct {
	Id      string      `bson:"_id"`
	Type    string      `bson:"type" json:"type"`
	NoteId  string      `bson:"noteId" json:"noteId"`
	Content []TextBlock `bson:"content" json:"content"`
}

type TextBlock struct {
	Style string `bson:"style" json:"style"`
	Text  string `bson:"text" json:"text"`
}

type Note struct {
	Id           string   `bson:"_id"`
	Title        string   `bson:"title"`
	CreationTime int64    `bson:"creation_time" json:"creation_time"`
	LastChange   int64    `bson:"last_change" json:"last_change"`
	Users        []string `bson:"users" json:"users"`
}
