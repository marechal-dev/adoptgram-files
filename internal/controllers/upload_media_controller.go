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

type UploadMediaController struct {
	uploader *s3manager.Uploader
}

func NewUploadMediaController(s3Session *session.Session) *UploadMediaController {
	uploader := s3manager.NewUploader(s3Session)

	return &UploadMediaController{
		uploader: uploader,
	}
}

func (umc *UploadMediaController) GetUploader() *s3manager.Uploader {
	return umc.uploader
}

func (umc *UploadMediaController) Handler(ctx *gin.Context) {
	form, _ := ctx.MultipartForm()
	file := form.File["image"][0]

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

	uploadError := services.UploadFileToR2(umc.uploader, buffer, fileIdentifier, fileType)

	if uploadError != nil {
		ctx.JSON(http.StatusBadRequest, responses.NewUploadFailedResponse(uploadError.Error()))
		return
	}

	successResponse := responses.NewMediaCreatedResponse(fileIdentifier)

	ctx.JSON(http.StatusCreated, successResponse)
}
