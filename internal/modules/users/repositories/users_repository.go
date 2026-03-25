package repositories

import "gorm.io/gorm"

type UsersRepository interface {}

type UsersRepositoryImpl struct {
    db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) UsersRepository {
    return &UsersRepositoryImpl{db: db}
}
