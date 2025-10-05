package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/zkfmapf123/fpg/internal"
	"go.uber.org/zap"
)

func main() {
	g := gin.Default()

	logger := internal.NewLogger()
	defer logger.Sync()

	g.GET("/", func(ctx *gin.Context) {
		log.Println("Hello, Default Server")
		logger.Info("msg", zap.String("message", "Hello, Default Server"))

		ctx.JSON(200, gin.H{
			"message": "Hello, Default Server",
		})
	})

	g.Run(":8080")
}
