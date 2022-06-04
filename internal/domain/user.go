package domain

import "time"

type User struct {
	Id          string
	DisplayName string
	Email       string
	Password    string
	DateOfBirth time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ChangePasswordRequest struct {
}

type ChangePasswordResponse struct {
}

type VerifyForgotPasswordRequest struct {
}

type VerifyForgotPasswordResponse struct {
}

type ForgotPasswordRequest struct {
}

type ForgotPasswordResponse struct {
}

type LoginRequest struct {
}

type LoginResponse struct {
}

type RegisterRequest struct {
}

type RegisterResponse struct {
}
