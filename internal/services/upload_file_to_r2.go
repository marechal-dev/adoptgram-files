package services

import (
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func UploadFileToR2(uploader *s3manager.Uploader, file io.Reader, key string, contentType string) error {
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(os.Getenv("AWS_BUCKET_NAME")),
		Key:         aws.String(key),
		ContentType: aws.String(contentType),
		Body:        file,
	})

	if err != nil {
		return err
	}

	return nil
}
