package repository_gorm

import (
	"nashtanet-backend-go/domain/entity"
	"nashtanet-backend-go/infrastructure/database"
)

type UserRepositoryGorm struct {
	db database.Gorm
}

const table = "users"

func NewUserRepositoryGorm(db database.Gorm) UserRepositoryGorm {
	return UserRepositoryGorm{db: db}
}

func (repo UserRepositoryGorm) GetAllUsers() ([]entity.User, error) {
	return nil, nil
}

func (repo UserRepositoryGorm) GetOneUser(user *entity.User) (*entity.User, error) {
	return nil, nil
}

func (repo UserRepositoryGorm) GetUserById(user *entity.User) (*entity.User, error) {
	return nil, nil
}

func (repo UserRepositoryGorm) CreateUser(user *entity.User) (*entity.User, error) {
	result, err := repo.db.Create(user, table)
	if err != nil {
		return nil, err
	}

	return result.(*entity.User), nil
}

func (repo UserRepositoryGorm) UpdateUser(user *entity.User) (*entity.User, error) {
	return nil, nil
}
