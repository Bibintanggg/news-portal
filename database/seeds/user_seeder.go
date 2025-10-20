package seeds

import (
	"bwa-news/internal/core/domain/model"
	"bwa-news/lib/conv"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) {
	bytes, err := conv.HashPassword("admin123")
	if err != nil {
		log.Fatal().Err(err).Msg("Error creating password hash")
	} // log fatal untuk langsung berhentiin program ketika ada err

	admin := model.User{
		Name:     "Admin",
		Email:    "admin@mail.com",
		Password: string(bytes),
	}

	if err := db.FirstOrCreate(&admin, model.User{Email: "admin@mail.com"}).Error; err != nil {
		log.Fatal().Err(err).Msg("Error seeding admin role")
	} else {
		log.Info().Msg("Admin role seeded successfully")
	}
}
