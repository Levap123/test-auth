package models

type User struct {
	Email    string `bson:"email" json:"email"`
	Salt     string `bson:"salt" json:"salt"`
	Password string `bson:"password" json:"password"`
}
