package jwt

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtService interface {
	Generate(user_id int) (string, error)
	Validate(token string) (int, error)
}

type jwtService struct {
}

func (js jwtService) Generate(user_id int) (string, error) {
	var jwtKey = []byte(os.Getenv("JWT_KEY"))
	myClaims := jwt.MapClaims{}
	myClaims["Id"] = user_id
	myClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	jwtInit := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims)

	token, err := jwtInit.SignedString(jwtKey)

	if err != nil {
		return "", err
	}
	return token, nil
}

func NewJwtService() jwtService {
	return jwtService{}
}

func (js jwtService) Validate(reqToken string) (int, error) {
	token, err := jwt.Parse(reqToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		return 0, err
	}

	claims, success := token.Claims.(jwt.MapClaims)
	if !success {
		return 0, nil
	}

	user_id, err := strconv.Atoi(fmt.Sprint(claims["Id"]))
	if err != nil {
		return 0, err
	}

	return user_id, nil
}
