package repository

import (
	"go-api/model"
)

type UsersRepository interface {
	Save(user model.Users)
	Update(user model.Users)
	Delete(userId int)
	FindById(userId int) (tags model.Users, err error)
	FindAll() []model.Users
}
