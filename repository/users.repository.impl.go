package repository

import (
	"errors"
	"go-api/data/request/users"
	"go-api/helper"
	"go-api/model"

	"gorm.io/gorm"
)

type UsersRepositoryImpl struct{
	Db *gorm.DB
}

func NewUsersRepoImpl(Db *gorm.DB) UsersRepository{
	return &UsersRepositoryImpl{Db: Db}

}