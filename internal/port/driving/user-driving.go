package driving

import (
	"context"

	"github.com/soap51/boss51.tech/internal/domain"
)

type IUserService interface {
	Register(context.Context, domain.RegisterRequest) (*domain.RegisterResponse, error)
	Login(context.Context, domain.LoginRequest) (*domain.LoginResponse, error)
	ForgotPassword(context.Context, domain.ForgotPasswordRequest) (*domain.ForgotPasswordResponse, error)
	VerifyForgotPassword(context.Context, domain.VerifyForgotPasswordRequest) (*domain.VerifyForgotPasswordResponse, error)
	ChangePassword(context.Context, domain.ChangePasswordRequest) (*domain.ChangePasswordResponse, error)
}
