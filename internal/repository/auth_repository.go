package repository

import (
	"context"
	"database/sql"

	"github.com/abu-umair/lms-be-microservice/internal/entity"
)

type IAuthRepository interface {
	GetUserByEmail(ctx context.Context, email string) (entity.Users, error) //?ctx: mengakses DB nya. User: entitynya (table di DB, terhubung di entity), (User, error) adalah return data atau error
}

type authRepository struct {
	db *sql.DB //?Menyimpan koneksi database
}

func (ar *authRepository) GetUserByEmail(ctx context.Context, email string) (entity.Users, error) {

}

func NewAuthRepository(db *sql.DB) IAuthRepository {
	return &authRepository{db: db}
}
