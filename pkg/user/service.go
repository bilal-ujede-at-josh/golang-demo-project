package user

import (
	"context"
	db "ispick-project-21022023/pkg/database"
	jwt "ispick-project-21022023/pkg/jwt"
	"ispick-project-21022023/pkg/model"
)

type UserService interface {
	FindUserByMobile(ctx context.Context, mobile string) (User, error)
	ValidateJwtToken(ctx context.Context, token string) (int, error)
	GetUserById(ctx context.Context, id int) (model.UserDetails, error)
}

type userService struct {
	db  db.DatabaseOps
	jwt jwt.JwtService
}

func NewUserService(db db.DatabaseOps, jwt jwt.JwtService) UserService {
	return &userService{
		db:  db,
		jwt: jwt,
	}
}

func (us userService) FindUserByMobile(ctx context.Context, mobile string) (User, error) {
	return User{}, nil
}

func (us userService) ValidateJwtToken(ctx context.Context, token string) (int, error) {
	return us.jwt.Validate(token)
}

func (us userService) GetUserById(ctx context.Context, id int) (model.UserDetails, error) {
	return us.db.FindUserById(ctx, id)
}
