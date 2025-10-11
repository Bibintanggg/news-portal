package config

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
}

func (cfg Config) ConnectionPostgres() (*Postgres, error) { // membuat package sendiri dengan  cfg/config
	dbConnString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.PsqlDB.User,
		cfg.PsqlDB.Password,
		cfg.PsqlDB.Host,
		cfg.PsqlDB.Port,
		cfg.PsqlDB.DBName,
	)

	db, err := gorm.Open(postgres.Open(dbConnString), &gorm.Config{}) // buat connect ke db
	if err != nil {
		log.Error().Err(err).Msg("[ConnectionPostgres-1] Failed to connect to database" + cfg.PsqlDB.Host)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Error().Err(err).Msg("[ConnecttionPostgres-2] Failed to connect to database" + cfg.PsqlDB.Host)
		return nil, err
	}

	sqlDB.SetMaxOpenConns(cfg.PsqlDB.DBMaxOpen)
	sqlDB.SetMaxIdleConns(cfg.PsqlDB.DBMaxIdle)

	return &Postgres{DB: db}, nil
}
