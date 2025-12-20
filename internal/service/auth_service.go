package service

import (
	"context"

	"github.com/abu-umair/lms-be-microservice/internal/repository"
	"github.com/abu-umair/lms-be-microservice/internal/utils"
	"github.com/abu-umair/lms-be-microservice/pb/auth"
)

type IAuthService interface {
	Register(ctx context.Context, request *auth.RegisterRequest) (*auth.RegisterResponse, error)
}

type authService struct {
	authRepository repository.IAuthRepository
}

func (as *authService) Register(ctx context.Context, request *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	//? ngecek email ke DB
	//* layer repository, utk akses DB (clean arsitektur)
	user, err := as.authRepository.GetUserByEmail(ctx, request.Email)
	if err != nil || user == nil {
		return nil, err
	}

	if user != nil {
		return &auth.RegisterResponse{
			Base: utils.BadRequestResponse("User already exist"),
		}, nil
	}

	//? jika emal sudah terdaftar/ada, di error in

	//? Hash password

	//? Insert ke DB
	return &auth.RegisterResponse{}, nil
}

func NewAuthService(authRepository repository.IAuthRepository) IAuthService {
	return &authService{
		authRepository: authRepository,
	}
}
