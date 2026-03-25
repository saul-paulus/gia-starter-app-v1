package services

import "gia-starter-app-V1/internal/modules/users/repositories"

type UsersService struct {
    repo repositories.UsersRepository
}

func NewUsersService(repo repositories.UsersRepository) *UsersService {
    return &UsersService{repo: repo}
}
