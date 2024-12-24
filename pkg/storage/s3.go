package storage

import (
	awsSDK "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var s3Client *s3.S3

func InitS3(region string) {
	sess := session.Must(session.NewSession(&awsSDK.Config{
		Region: awsSDK.String(region),
	}))
	s3Client = s3.New(sess)
}

func GetS3Client() *s3.S3 {
	return s3Client
}