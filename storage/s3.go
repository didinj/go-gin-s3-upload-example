package storage

import (
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/yourusername/go-upload-api/config"
)

var s3Client *s3.Client

func init() {
	if config.StorageType == "s3" {
		cfg, err := awsConfig.LoadDefaultConfig(context.TODO())
		if err != nil {
			panic(err)
		}
		s3Client = s3.NewFromConfig(cfg)
		UploadFile = uploadS3
	}
}

func uploadS3(filename string, file io.Reader) (string, error) {
	_, err := s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(config.S3Bucket),
		Key:    aws.String(filename),
		Body:   file,
		ACL:    "public-read",
	})
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", config.S3Bucket, config.S3Region, filename), nil
}
