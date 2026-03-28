package repositories

import (
	"database/sql"
	"regexp"
	"testing"
	"time"

	"gia-starter-app-V1/internal/modules/users/domain"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// setupMockDB membuat instance GORM dengan sqlmock sebagai driver.
func setupMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, *sql.DB) {
	t.Helper()

	sqlDB, mock, err := sqlmock.New()
	require.NoError(t, err, "gagal membuat sqlmock")

	dialector := postgres.New(postgres.Config{
		Conn: sqlDB,
	})

	db, err := gorm.Open(dialector, &gorm.Config{})
	require.NoError(t, err, "gagal membuka koneksi GORM dengan sqlmock")

	return db, mock, sqlDB
}

// -----------------------------------------------------------------------
// CreateUser
// -----------------------------------------------------------------------

func TestCreateUser(t *testing.T) {
	tests := []struct {
		name    string
		user    *domain.Users
		setup   func(mock sqlmock.Sqlmock)
		wantErr bool
	}{
		{
			name: "success - user berhasil diinsert",
			user: &domain.Users{
				Username: "saul",
				Email:    "saul@mail.com",
				Password: "hashed_password",
				RoleID:   1,
			},
			// GORM postgres menggunakan RETURNING "id", tanpa mengirim ID sebagai arg
			// Query: INSERT INTO "users" (username,email,password,id_role,is_active,created_at,updated_at) VALUES ($1...$7) RETURNING "id"
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users"`)).
					WithArgs(
						"saul",          // username
						"saul@mail.com", // email
						"hashed_password", // password
						1,               // id_role
						false,           // is_active
						sqlmock.AnyArg(), // created_at
						sqlmock.AnyArg(), // updated_at
					).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
				mock.ExpectCommit()
			},
			wantErr: false,
		},
		{
			name: "error - insert gagal karena DB error",
			user: &domain.Users{
				Username: "saul",
				Email:    "saul@mail.com",
				Password: "hashed_password",
				RoleID:   1,
			},
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users"`)).
					WillReturnError(sql.ErrConnDone)
				mock.ExpectRollback()
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, sqlDB := setupMockDB(t)
			defer sqlDB.Close()

			tc.setup(mock)

			repo := NewUsersRepository(db)
			err := repo.CreateUser(tc.user)

			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.NoError(t, mock.ExpectationsWereMet(), "ada ekspektasi sqlmock yang tidak terpenuhi")
		})
	}
}

// -----------------------------------------------------------------------
// FindByEmailUser
// -----------------------------------------------------------------------

func TestFindByEmailUser(t *testing.T) {
	now := time.Now()

	// Query aktual GORM:
	// SELECT * FROM "users" WHERE email =$1 ORDER BY "users"."id" LIMIT $2
	const findQuery = `SELECT * FROM "users" WHERE email =$1 ORDER BY "users"."id" LIMIT $2`

	tests := []struct {
		name     string
		email    string
		setup    func(mock sqlmock.Sqlmock)
		wantUser *domain.Users
		wantErr  bool
	}{
		{
			name:  "success - user ditemukan berdasarkan email",
			email: "saul@mail.com",
			setup: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "username", "email", "password", "id_role", "is_active", "created_at", "updated_at"}).
					AddRow(1, "saul", "saul@mail.com", "hashed_pwd", 1, true, now, now)
				mock.ExpectQuery(regexp.QuoteMeta(findQuery)).
					WithArgs("saul@mail.com", 1). // email + LIMIT 1
					WillReturnRows(rows)
			},
			wantUser: &domain.Users{
				ID:       1,
				Username: "saul",
				Email:    "saul@mail.com",
				RoleID:   1,
			},
			wantErr: false,
		},
		{
			name:  "error - user tidak ditemukan (record not found)",
			email: "notfound@mail.com",
			setup: func(mock sqlmock.Sqlmock) {
				// Return empty rows -> GORM akan kembalikan gorm.ErrRecordNotFound
				rows := sqlmock.NewRows([]string{"id", "username", "email", "password", "id_role", "is_active", "created_at", "updated_at"})
				mock.ExpectQuery(regexp.QuoteMeta(findQuery)).
					WithArgs("notfound@mail.com", 1).
					WillReturnRows(rows)
			},
			wantErr: true, // gorm.ErrRecordNotFound
		},
		{
			name:  "error - query gagal karena DB error",
			email: "saul@mail.com",
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(findQuery)).
					WithArgs("saul@mail.com", 1).
					WillReturnError(sql.ErrConnDone)
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, sqlDB := setupMockDB(t)
			defer sqlDB.Close()

			tc.setup(mock)

			repo := NewUsersRepository(db)
			user, err := repo.FindByEmailUser(tc.email)

			if tc.wantErr {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, user)
				assert.Equal(t, tc.wantUser.ID, user.ID)
				assert.Equal(t, tc.wantUser.Email, user.Email)
				assert.Equal(t, tc.wantUser.Username, user.Username)
				assert.Equal(t, tc.wantUser.RoleID, user.RoleID)
			}

			assert.NoError(t, mock.ExpectationsWereMet(), "ada ekspektasi sqlmock yang tidak terpenuhi")
		})
	}
}
