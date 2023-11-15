package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/marechal-dev/adoptgram-files/internal/controllers"
)

func main() {
	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		log.Fatal("Could not find \"PORT\" env variable")
	}

	serverAddr := fmt.Sprintf("0.0.0.0:%d", port)

	router := gin.Default()

	router.MaxMultipartMemory = 64 << 20 // 64 MiB

	uploadMediasController := controllers.NewUploadMediasController()

	router.POST("/medias/upload", uploadMediasController.Handler)

	router.Run(serverAddr)
}
