package services

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func UploadFileToR2(uploader *s3manager.Uploader, file io.Reader, key string, contentType string) {
	uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String("a"),
		Key:         aws.String(key),
		ContentType: aws.String(contentType),
		Body:        file,
	})
}
