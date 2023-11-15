package controllers

import (
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
	"github.com/marechal-dev/adoptgram-files/internal/responses"
	"github.com/marechal-dev/adoptgram-files/internal/services"
)

type UploadMediasController struct {
	uploader *s3manager.Uploader
}

func NewUploadMediasController() *UploadMediasController {
	session, err := session.NewSession(&aws.Config{
		Region: aws.String("auto"),
	})

	if err != nil {
		log.Fatal("Could not create an AWS Session")
	}

	uploader := s3manager.NewUploader(session)

	return &UploadMediasController{
		uploader: uploader,
	}
}

func (umc *UploadMediasController) GetUploader() *s3manager.Uploader {
	return umc.uploader
}

func (umc *UploadMediasController) Handler(ctx *gin.Context) {
	form, _ := ctx.MultipartForm()
	files := form.File["files[]"]
	filesUrls := []string{}

	for _, file := range files {
		buffer, err := file.Open()
		defer buffer.Close()

		if err != nil {
			ctx.JSON(http.StatusBadRequest, responses.NewCouldNotParseFileResponse())
			return
		}

		fileKey, err := services.GetFileKeyService(file.Filename)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, responses.NewCouldNotParseFileResponse())
			return
		}

		go services.UploadFileToR2(umc.uploader, buffer, fileKey, "image/png")

		filesUrls = append(filesUrls, fileKey)
	}

	successResponse := responses.NewMediasCreatedResponse(filesUrls)

	ctx.JSON(http.StatusCreated, successResponse)
}
