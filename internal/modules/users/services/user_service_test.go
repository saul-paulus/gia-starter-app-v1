package services

import (
	"testing"

	"gia-starter-app-V1/internal/modules/users/domain"
	"gia-starter-app-V1/internal/modules/users/dto"
	"gia-starter-app-V1/internal/modules/users/repositories/mocks"
	appErrors "gia-starter-app-V1/internal/shared/errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	tests := []struct {
		name          string
		mockFindEmail func(email string) (*domain.Users, error)
		mockCreate    func(user *domain.Users) error
		req           dto.CreateUser
		wantErr       bool
		wantErrCode   string
	}{
		{
			name: "success - user baru berhasil dibuat",
			// Email belum ada -> gorm.ErrRecordNotFound
			mockFindEmail: func(email string) (*domain.Users, error) {
				return nil, gorm.ErrRecordNotFound
			},
			mockCreate: func(user *domain.Users) error {
				return nil
			},
			req: dto.CreateUser{
				Username: "saul",
				Email:    "saul@mail.com",
				RoleID:   1,
				Password: "password123",
			},
			wantErr: false,
		},
		{
			name: "error - email sudah terdaftar",
			// Email ditemukan -> return user dengan ID terisi
			mockFindEmail: func(email string) (*domain.Users, error) {
				return &domain.Users{ID: 1}, nil
			},
			req: dto.CreateUser{
				Username: "saul",
				Email:    "saul@mail.com",
				RoleID:   1,
				Password: "password123",
			},
			wantErr:     true,
			wantErrCode: "EMAIL_EXISTS",
		},
		{
			name: "error - gagal query email ke database",
			// DB error bukan record not found -> harus return ErrInternal
			mockFindEmail: func(email string) (*domain.Users, error) {
				return nil, gorm.ErrInvalidDB
			},
			req: dto.CreateUser{
				Username: "saul",
				Email:    "saul@mail.com",
				RoleID:   1,
				Password: "password123",
			},
			wantErr:     true,
			wantErrCode: appErrors.ErrInternal.Code,
		},
		{
			name: "error - gagal insert user ke database",
			// Email belum ada, tapi create gagal
			mockFindEmail: func(email string) (*domain.Users, error) {
				return nil, gorm.ErrRecordNotFound
			},
			mockCreate: func(user *domain.Users) error {
				return gorm.ErrInvalidDB
			},
			req: dto.CreateUser{
				Username: "saul",
				Email:    "saul@mail.com",
				RoleID:   1,
				Password: "password123",
			},
			wantErr:     true,
			wantErrCode: appErrors.ErrInternal.Code,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := &mocks.UsersRepositoryMock{
				FindByEmailUserFunc: tc.mockFindEmail,
				CreateUserFunc:      tc.mockCreate,
			}

			svc := NewUsersService(mockRepo)
			err := svc.CreateUser(tc.req)

			if tc.wantErr {
				require.Error(t, err)
				appErr, ok := err.(*appErrors.AppError)
				require.True(t, ok, "error harus bertipe *AppError, bukan: %T", err)
				assert.Equal(t, tc.wantErrCode, appErr.Code)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
