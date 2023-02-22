package auth

import (
	"context"
	db "ispick-project-21022023/pkg/database"
	jwt "ispick-project-21022023/pkg/jwt"
	"log"
)

type AuthService interface {
	SendOtp(ctx context.Context, mobile string) error
	VerifyOtp(ctx context.Context, mobile string, otp int) error
	JwtToken(ctx context.Context, mobile string) (string, error)
}

type authService struct {
	db  db.DatabaseOps
	jwt jwt.JwtService
}

func (as authService) SendOtp(ctx context.Context, mobile string) error {

	return as.db.InsertOtp(ctx, mobile)
}

func (as authService) VerifyOtp(ctx context.Context, mobile string, otp int) error {
	err := as.db.ValidateOtp(ctx, mobile, otp)
	if err != nil {
		log.Println("Error validationg OTP", err)
		return err
	}
	return nil
}

func (as authService) JwtToken(ctx context.Context, mobile string) (string, error) {
	userId, err := as.db.FindUserByMobile(ctx, mobile)
	if err != nil {
		return "", err
	}
	token, err := as.jwt.Generate(userId)
	if err != nil {
		return "", err
	}
	return token, nil
}

func NewAuthService(db db.DatabaseOps, jwt jwt.JwtService) AuthService {
	return &authService{
		db:  db,
		jwt: jwt,
	}
}
