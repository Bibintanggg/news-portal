package repository // berfungsi akses ke DB (CRUD) dan komunikasi dengan service

import (
	"bwa-news/internal/core/domain/entity"
	"bwa-news/internal/core/domain/model"
	"context"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

var err error
var code string

type AuthRepository interface { // global pkg
	GetUserByEmail(ctx context.Context, req entity.LoginRequest) (*entity.UserEntity, error)
}

type authRepository struct { // private implement
	db *gorm.DB
}

// GetUserByEmail implements AuthRepo
func (a *authRepository) GetUserByEmail(ctx context.Context, req entity.LoginRequest) (*entity.UserEntity, error) {
	var modelUser model.User

	err := a.db.Where("email = ?", req.Email).First(&modelUser).Error
	if err != nil {
		code = "[REPOSITORY] GetUserByEmail - 1"
		log.Errorw(code, err)
		return nil, err
	}

	resp := entity.UserEntity{
		ID:       modelUser.ID,
		Name:     modelUser.Name,
		Email:    modelUser.Email,
		Password: modelUser.Password,
	}

	return &resp, nil

	panic("Unimplemented")
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}
