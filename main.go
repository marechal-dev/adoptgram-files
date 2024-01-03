package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/marechal-dev/adoptgram-files/internal/controllers"
)

func main() {
	err := godotenv.Load(".env")

	appEnv := os.Getenv("APP_ENV")

	if appEnv == "development" {
		os.Setenv("AWS_BUCKET_NAME", "adoptgram-files-test")
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		log.Fatal("Could not find/parse \"PORT\" env variable")
	}

	serverAddr := fmt.Sprintf("127.0.0.1:%d", port)

	router := gin.Default()

	router.MaxMultipartMemory = 10 << 20 // 64 MiB

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")

	session, err := session.NewSession(&aws.Config{
		Region:   aws.String("auto"),
		Endpoint: aws.String(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountID)),
	})

	if err != nil {
		log.Fatal("Could not create an AWS Session")
	}

	uploadMediasController := controllers.NewUploadMediasController(session)
	uploadMediaController := controllers.NewUploadMediaController(session)

	router.POST("/api/media/upload", uploadMediaController.Handler)
	router.POST("/api/media/bulk-upload", uploadMediasController.Handler)

	router.Run(serverAddr)
}
