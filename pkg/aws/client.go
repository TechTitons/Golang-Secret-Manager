package aws

import (
	"context"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"

	"github.com/TechTitons/Golang-Secret-Manager/config"
)

func NewSecretManagerClient(cfg *config.Config) (*secretsmanager.Client, error) {

	awsCfg, err := awsconfig.LoadDefaultConfig(
		context.TODO(),
		awsconfig.WithRegion(cfg.AWSRegion),
	)
	if err != nil {
		return nil, err
	}

	return secretsmanager.NewFromConfig(awsCfg), nil
}