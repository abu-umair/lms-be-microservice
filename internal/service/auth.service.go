package service

import (
	"context"

	"github.com/abu-umair/lms-be-microservice/pb/auth"
)

type IAuthService interface {
	Register(ctx context.Context, request *auth.RegisterRequest) (*auth.RegisterResponse, error)
}

type authService struct{}

func (as *authService) Register(ctx context.Context, request *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	//? ngecek email ke DB

	//? jika emal sudah terdaftar/ada, di error in

	//? Hash password

	//? Insert ke DB 
	return &auth.RegisterResponse{}, nil
}

func NewAuthService() IAuthService {
	return &authService{}
}
