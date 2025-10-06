package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zkfmapf123/fpg/internal"
)

func main() {
	g := gin.Default()

	logger := internal.NewLogger()
	defer logger.Sync()

	g.GET("/", func(ctx *gin.Context) {

		internal.InfoLogger("home", map[string]any{
			"age":  "32",
			"name": "leedonggyu",
			"job":  "devops",
		})

		internal.WarnLogger("home", map[string]any{
			"age":  "32",
			"name": "leedonggyu",
			"job":  "devops",
		})

		ctx.JSON(200, gin.H{
			"message": "Hello, Default Server",
		})
	})

	go looplogger()
	g.Run(":8080")
}

func looplogger() {
	i := 1
	for {
		internal.InfoLogger("calculator", map[string]any{
			"result":   i,
			"madeby":   "leedonggyu",
			"result*2": i * i,
		})

		i++

		time.Sleep(time.Second * 2)
	}
}
