package models

type User struct {
	Email    string `bson:"email"`
	Salt     string `bson:"salt"`
	Password string `bson:"password"`
}
