package auth

import (
	"bwa-news/config"
	"bwa-news/internal/core/domain/entity"
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
	data.RegisteredClaims.Issuer = o.issuer                         // memberi tahu siapa pembuat
	data.RegisteredClaims.NotBefore = jwt.NewNumericDate(now)       // token ini gaboleh dipakai sebelum (now)

	acToken := jwt.NewWithClaims(jwt.SigningMethodHS256, data) // bikin tokennya

	accessToken, err := acToken.SignedString([]byte(o.signingKey)) //tandatangan token secret key
	if err != nil {
		return "", 0, err
	}
	return accessToken, expiresAt.Unix(), nil
}

func (o *Options) VerifyAccessToken(token string) (*entity.JwtData, error) {
	panic("Unimplemented")
}

func NewJwt(cfg *config.Config) Jwt {
	opt := new(Options)
	opt.signingKey = cfg.App.JWTSecretKey
	opt.issuer = cfg.App.JWTIssuer

	return opt
}
