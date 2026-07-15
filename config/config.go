package config

import "os"

type Config struct {
	AWSAccessKey string
	AWSSecretKey string
	AWSRegion    string
	Port         string
}

func LoadConfig() *Config {
	return &Config{
		AWSAccessKey: os.Getenv("AWS_ACCESS_KEY_ID"),
		AWSSecretKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
		AWSRegion:    os.Getenv("AWS_REGION"),
		Port:         os.Getenv("PORT"),
	}
}
