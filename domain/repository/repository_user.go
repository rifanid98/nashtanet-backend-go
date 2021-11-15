package repository

import "nashtanet-backend-go/domain/entity"

type UserRepository interface {
	GetAllUsers() ([]entity.User, error)
	GetOneUser(user *entity.User) (*entity.User, error)
	GetUserById(user *entity.User) (*entity.User, error)
	CreateUser(user *entity.User) (*entity.User, error)
	UpdateUser(user *entity.User) (*entity.User, error)
}
