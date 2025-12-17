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
	return &auth.RegisterResponse{}, nil
}

func NewAuthService() IAuthService {
	return &authService{}
}
