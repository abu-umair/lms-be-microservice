package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/abu-umair/lms-be-microservice/internal/entity"
)

type IAuthRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*entity.Users, error) //?ctx: mengakses DB nya. User: entitynya (table di DB, terhubung di entity), (User, error) adalah return data atau error
	InsertUser(ctx context.Context, user *entity.Users) error
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

func (ar *authRepository) InsertUser(ctx context.Context, user *entity.Users) error {
	_, err := ar.db.ExecContext(
		ctx,
		`INSERT INTO users (id, full_name, email, password, role_code, created_at, created_by, updated_at, updated_by, deleted_at,deleted_by, is_deleted)
                VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`,

		user.Id,
		user.FullName,
		user.Email,
		user.Password,
		user.RoleCode,
		user.CreatedAt,
		user.CreatedBy,
		user.UpdatedAt,
		user.UpdatedBy,
		user.DeletedAt,
		user.DeletedBy,
		user.IsDeleted,
	)

	if err != nil {
		return err
	}

	return nil
}

func NewAuthRepository(db *sql.DB) IAuthRepository {
	return &authRepository{db: db}
}
