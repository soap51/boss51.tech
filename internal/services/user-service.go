package services

import (
	"context"

	"github.com/soap51/boss51.tech/internal/domain"
	"github.com/soap51/boss51.tech/internal/port/driven"
)

type userService struct {
	userRepository driven.IUserRepository
}

func NewUserService(userRepository driven.IUserRepository) *userService {
	return &userService{
		userRepository: userRepository,
	}
}

func (svc *userService) Register(context.Context, domain.RegisterRequest) (*domain.RegisterResponse, error) {
	return nil, nil
}

func (svc *userService) Login(context.Context, domain.LoginRequest) (*domain.LoginResponse, error) {
	return nil, nil
}

func (svc *userService) ForgotPassword(context.Context, domain.ForgotPasswordRequest) (*domain.ForgotPasswordResponse, error) {
	return nil, nil
}

func (svc *userService) VerifyForgotPassword(context.Context, domain.VerifyForgotPasswordRequest) (*domain.VerifyForgotPasswordResponse, error) {
	return nil, nil
}

func (svc *userService) ChangePassword(context.Context, domain.ChangePasswordRequest) (*domain.ChangePasswordResponse, error) {
	return nil, nil
}
