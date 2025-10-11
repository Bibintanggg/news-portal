package app

import (
	"bwa-news/config"

	"github.com/rs/zerolog/log"
)

func RunServer() {
	cfg := config.NewConfig()
	_, err := cfg.ConnectionPostgres()
	if err != nil {
		log.Fatal().Msgf("Error to connect to database: %v", err)
	}
}
