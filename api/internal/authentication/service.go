package authentication

import (
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/infinitemax/bookAgain/internal/users"
	"os"
	"time"
)

type IAuthenticationService interface {
	CreateUser(ctx context.Context, user *users.User) error
}

type Service struct {
	db IAuthenticationService
}

type UserClaims struct {
	Id       uuid.UUID `json:"id"`
	UserName string    `json:"user_name"`
	*jwt.StandardClaims
}

type AccessToken struct {
	token string
}

func NewService(db IAuthenticationService) *Service {
	return &Service{db: db}
}

func NewAccessToken(claims UserClaims) (*AccessToken, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	newToken := &AccessToken{}
	var err error
	newToken.token, err = accessToken.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	if err != nil {
		return nil, err
	}

	return newToken, nil
}

func NewRefreshToken(claims jwt.StandardClaims) (*AccessToken, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	newToken := &AccessToken{}
	var err error
	newToken.token, err = refreshToken.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	if err != nil {
		return nil, err
	}

	return newToken, nil

}

func ParseAccessToken(accessToken string) *UserClaims {
	parsedAccessToken, _ := jwt.ParseWithClaims(accessToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})

	return parsedAccessToken.Claims.(*UserClaims)
}

func ParseRefreshToken(refreshToken string) *jwt.StandardClaims {
	parsedRefreshToken, _ := jwt.ParseWithClaims(refreshToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})

	return parsedRefreshToken.Claims.(*jwt.StandardClaims)
}

func (s *Service) CreateUser(ctx context.Context, user *users.User) error {

	user.CreatedAt = time.Now()
	// hash password!

	err := s.db.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
