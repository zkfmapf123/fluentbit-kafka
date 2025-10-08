package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zkfmapf123/fpg/internal"
)

func main() {
	g := gin.Default()

	// logger
	logger := internal.NewLogger()
	defer logger.Sync()

	// kafak
	kafka := internal.MustNewPubsub(logger)

	g.GET("/", func(ctx *gin.Context) {

		logger.InfoLogger("home", map[string]any{
			"age":  "32",
			"name": "leedonggyu",
			"job":  "devops",
		})

		logger.ErrorLogger("home", map[string]any{
			"age":  "32",
			"name": "leedonggyu",
			"job":  "devops",
		})

		ctx.JSON(200, gin.H{
			"message": "Hello, Default Server",
		})
	})

	go looplogger(kafka)
	g.Run(":8080")
}

func looplogger(kafka *internal.Pubsub) {
	i := 1
	for {

		kafka.Producer("home", map[string]any{
			"age":  "32",
			"name": "leedonggyu",
			"job":  "devops",
		})

		// internal.InfoLogger("calculator", map[string]any{
		// 	"result":   i,
		// 	"madeby":   "leedonggyu",
		// 	"result*2": i * i,
		// })

		i++

		time.Sleep(time.Second * 2)
	}
}
