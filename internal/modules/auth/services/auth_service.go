package services

import "gia-starter-app-V1/internal/modules/auth/repositories"

type AuthService struct {
    repo repositories.AuthRepository
}

func NewAuthService(repo repositories.AuthRepository) *AuthService {
    return &AuthService{repo: repo}
}
