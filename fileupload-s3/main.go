package main

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Region          string `envconfig:"AWS_REGION"`
	Endpoint        string `envconfig:"AWS_ENDPOINT"`
	AccessKeyID     string `envconfig:"AWS_ACCESS_KEY_ID"`
	SecretAccessKey string `envconfig:"AWS_SECRET_ACCESS_KEY"`
	Environment     string `envconfig:"ENVIRONMENT"`
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	var c Config
	envconfig.Process("", &c)
	cfg := aws.Config{
		Region: c.Region,
		Credentials: aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
			return aws.Credentials{
				AccessKeyID:     c.AccessKeyID,
				SecretAccessKey: c.SecretAccessKey,
			}, nil
		}),
	}

	if c.Environment == "dev" {
		cfg.EndpointResolverWithOptions = aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL: c.Endpoint,
			}, nil
		})
	}
	s3Client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})
	bucketBasics := BucketBasics{S3Client: s3Client}
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	bucketBasics.UploadFile(ctx, "test-bucket", "dummy.json", "./dummy.json")
	defer cancel()
}
