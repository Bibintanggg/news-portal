package service

import (
	"bwa-news/config"
	"bwa-news/internal/adapter/repository"
	"bwa-news/internal/core/domain/entity"
	"bwa-news/lib/auth"
	"bwa-news/lib/conv"
	"context"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
)

var err string
var code string

type AuthService interface {
	GetUserByEmail(ctx context.Context, req entity.LoginRequest) (*entity.AccessToken, error)
}

type authService struct {
	authRepository repository.AuthRepository
	cfg            *config.Config // untuk generate token
	jwtToken       auth.Jwt       // untuk generate token
}

func (a *authService) GetUserByEmail(ctx context.Context, req entity.LoginRequest) (*entity.AccessToken, error) {
	result, err := a.authRepository.GetUserByEmail(ctx, req)
	if err != nil {
		code = "[Service] GetUserByEmail - 1"
		log.Errorw(code, err)
		return nil, err
	}

	if checkPass := conv.CheckPasswordHash(req.Password, result.Password); !checkPass {
		code = "[Service] GetUserByEmail - 2"
		log.Errorw(code, "Invalid Password")
		return nil, err
	}

	jwtData := entity.JwtData{
		UserId: float64(result.ID),
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)), // masa berlakunya jwt
			ID:        string(result.ID),
		},
	}

	accesToken, expiresAt, err := a.jwtToken.GenerateToken(&jwtData)
	if err != nil {
		code = "[Service] GetUserByEmail - 3"
		log.Errorw(code, "Invalid Password")
		return nil, err
	}

	resp := entity.AccessToken{
		AccessToken: accesToken,
		ExpiresAt:   expiresAt,
	}

	return &resp, nil 
}

func NewAuthService(authRepository repository.AuthRepository, cfg *config.Config, jwtToken auth.Jwt) AuthService {
	return &authService{ // untuk generate token JWT
		authRepository: authRepository,
		cfg:            cfg,
		jwtToken:       jwtToken,
	}
}
