package repositories

import "gorm.io/gorm"

type AuthRepository interface {}

type AuthRepositoryImpl struct {
    db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
    return &AuthRepositoryImpl{db: db}
}
