package config

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config" // membuat var baru supaya ga confl
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/rs/zerolog/log"
)

func (cfg Config) LoadAwsConfig() aws.Config {
	conf, err := awsConfig.LoadDefaultConfig(context.TODO(),
		awsConfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			cfg.R2.ApiKey, cfg.R2.ApiSecret, "",
		)), awsConfig.WithRegion("auto"))
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to load aws config")
	}

	log.Info().Msg("Success! Loaded AWS Config")

	return conf
}
