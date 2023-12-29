package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
	"github.com/marechal-dev/adoptgram-files/internal/responses"
	"github.com/marechal-dev/adoptgram-files/internal/services"
)

type UploadMediasController struct {
	uploader *s3manager.Uploader
}

func NewUploadMediasController(s3Session *session.Session) *UploadMediasController {
	uploader := s3manager.NewUploader(s3Session)

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

		fileIdentifier, err := services.GetFileKeyService(file.Filename)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, responses.NewCouldNotParseFileResponse())
			return
		}

		fileExtension := strings.Split(file.Filename, ".")[1]
		fileType := fmt.Sprintf("image/%s", fileExtension)

		go services.UploadFileToR2(umc.uploader, buffer, fileIdentifier, fileType)

		filesUrls = append(filesUrls, fileIdentifier)
	}

	successResponse := responses.NewMediasCreatedResponse(filesUrls)

	ctx.JSON(http.StatusCreated, successResponse)
}
