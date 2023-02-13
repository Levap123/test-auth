package repository

type Repository struct {
	UserRepo
}

type UserRepo interface {
	Create(email, password, salt string) error
}
