package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	cfg := aws.Config{
		Region: "us-west-2",
		EndpointResolverWithOptions: aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL: "http://localhost:9000",
			}, nil
		}),
		Credentials: aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
			return aws.Credentials{
				AccessKeyID:     "suryansh",
				SecretAccessKey: "suryansh",
			}, nil
		}),
	}
	s3Client := s3.NewFromConfig(cfg)
	bucketBasics := BucketBasics{S3Client: s3Client}
	context := context.Background()
	bucketBasics.UploadFile(context, "test-bucket", "dummy.json", "./dummy.json")
}
