package mocks

import "gia-starter-app-V1/internal/modules/users/domain"

// UsersRepositoryMock adalah implementasi mock dari UsersRepository untuk keperluan unit testing.
// Setiap field adalah fungsi yang bisa di-set per test case.
type UsersRepositoryMock struct {
	CreateUserFunc      func(user *domain.Users) error
	FindByEmailUserFunc func(email string) (*domain.Users, error)
}

// CreateUser memanggil CreateUserFunc jika di-set, atau mengembalikan nil secara default.
func (m *UsersRepositoryMock) CreateUser(user *domain.Users) error {
	if m.CreateUserFunc != nil {
		return m.CreateUserFunc(user)
	}
	return nil
}

// FindByEmailUser memanggil FindByEmailUserFunc jika di-set, atau mengembalikan nil secara default.
func (m *UsersRepositoryMock) FindByEmailUser(email string) (*domain.Users, error) {
	if m.FindByEmailUserFunc != nil {
		return m.FindByEmailUserFunc(email)
	}
	return nil, nil
}
