package service

import (
	"context"
	"time"

	"github.com/abu-umair/lms-be-microservice/internal/entity"
	"github.com/abu-umair/lms-be-microservice/internal/repository"
	"github.com/abu-umair/lms-be-microservice/internal/utils"
	"github.com/abu-umair/lms-be-microservice/pb/auth"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Register(ctx context.Context, request *auth.RegisterRequest) (*auth.RegisterResponse, error)
}

type authService struct {
	authRepository repository.IAuthRepository
}

func (as *authService) Register(ctx context.Context, request *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	//? apakah password sama dengan confirm password
	if request.Password != request.PasswordConfirmation {
		return &auth.RegisterResponse{
			Base: utils.BadRequestResponse("Password is not matched"),
		}, nil
	}
	
	//? ngecek email ke DB
	//* layer repository, utk akses DB (clean arsitektur)
	user, err := as.authRepository.GetUserByEmail(ctx, request.Email)
	if err != nil || user == nil {
		return nil, err
	}

	//* jika emal sudah terdaftar/ada, di error in
	if user != nil {
		return &auth.RegisterResponse{
			Base: utils.BadRequestResponse("User already exist"),
		}, nil
	}

	//? Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		return nil, err
	}

	//? Insert ke DB
	newUser := entity.Users{
		Id:        uuid.NewString(),
		FullName:  request.FullName,
		Email:     request.Email,
		Password:  string(hashedPassword),
		RoleCode:  entity.UserRoleUser,
		CreatedAt: time.Now(),
		CreatedBy: &request.FullName,
	}

	err = as.authRepository.InsertUser(ctx, &newUser)
	if err != nil {
		return nil, err
	}

	return &auth.RegisterResponse{
		Base: utils.SuccessResponse("User is registered"),
	}, nil
}

func NewAuthService(authRepository repository.IAuthRepository) IAuthService {
	return &authService{
		authRepository: authRepository,
	}
}
