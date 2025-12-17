package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/abu-umair/lms-be-microservice/internal/entity"
)

type IAuthRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*entity.Users, error) //?ctx: mengakses DB nya. User: entitynya (table di DB, terhubung di entity), (User, error) adalah return data atau error
}

type authRepository struct {
	db *sql.DB //?Menyimpan koneksi database
}

func (ar *authRepository) GetUserByEmail(ctx context.Context, email string) (*entity.Users, error) {
	row := ar.db.QueryRowContext(ctx, "SELECT id, email, password, full_name FROM users WHERE email = $1 AND is_deleted = false", email)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var user entity.Users
	err := row.Scan(
		&user.Id,
		&user.Email,
		&user.Password,
		&user.FullName,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}


func NewAuthRepository(db *sql.DB) IAuthRepository {
	return &authRepository{db: db}
}
