package auth

import (
	"bwa-news/config"
	"bwa-news/internal/core/domain/entity"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Jwt interface {
	GenerateToken(data *entity.JwtData) (string, int64, error)
	VerifyAccessToken(token string) (*entity.JwtData, error)
}

type Options struct {
	signingKey string // membuat dan verif tanda token
	issuer     string // menanda siapa yg membuat jwt tsb
}

func (o *Options) GenerateToken(data *entity.JwtData) (string, int64, error) {
	now := time.Now().Local()            //mengatur waktu dan membuat waktu timestamp sekarang
	expiresAt := now.Add(time.Hour * 24) // token berlaku selama 24jam

	data.RegisteredClaims.ExpiresAt = jwt.NewNumericDate(expiresAt) // token ini expired pada kapan
	data.RegisteredClaims.Issuer = o.issuer                         // token ini dibuat oleh siapa
	data.RegisteredClaims.NotBefore = jwt.NewNumericDate(now)       // token ini bisa dipakai sebelum (now)

	acToken := jwt.NewWithClaims(jwt.SigningMethodHS256, data)

	accessToken, err := acToken.SignedString([]byte(o.signingKey))
	if err != nil {
		return "", 0, err
	}
	return accessToken, expiresAt.Unix(), nil
}

func (o *Options) VerifyAccessToken(token string) (*entity.JwtData, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		}
		return []byte(o.signingKey), nil
	})

	if err != nil {
		return nil, err
	}

	if parsedToken.Valid {
		claim, ok := parsedToken.Claims.(jwt.MapClaims)
		if !ok || parsedToken.Valid {
			return nil, err
		}

		jwtData := &entity.JwtData{
			UserId: claim["user_id"].(float64),
		}

		return jwtData, nil
	}

	return nil, fmt.Errorf("token is invalid")
}

func NewJwt(cfg *config.Config) Jwt {
	opt := new(Options)
	opt.signingKey = cfg.App.JWTSecretKey
	opt.issuer = cfg.App.JWTIssuer

	return opt
}
