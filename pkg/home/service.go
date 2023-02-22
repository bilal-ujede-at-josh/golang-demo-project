package home

import (
	"context"
	"ispick-project-21022023/pkg/database"
	"ispick-project-21022023/pkg/jwt"
	"ispick-project-21022023/pkg/model"
)

type HomeService interface {
	// HomeContent(user_id int) (model.AppHomeScreen, error)
	AppVersion(ctx context.Context) (model.AppVersion, error)
}

type homeService struct {
	db  database.DatabaseOps
	jwt jwt.JwtService
}

// func (hs homeService) HomeContent(user_id int) (model.AppHomeScreen, error) {
// 	return appHomeScreen, nil
// }

func NewHomeService(db database.DatabaseOps, jwt jwt.JwtService) HomeService {
	return &homeService{
		db:  db,
		jwt: jwt,
	}
}

func (hs homeService) AppVersion(ctx context.Context) (model.AppVersion, error) {
	return hs.db.AppVersion(ctx)
}
