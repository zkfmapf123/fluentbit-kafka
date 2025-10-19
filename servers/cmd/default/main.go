package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zkfmapf123/fpg/config"
	"github.com/zkfmapf123/fpg/internal"
	"github.com/zkfmapf123/fpg/models"
	"github.com/zkfmapf123/fpg/repository"
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

	db := config.NewDBConn()
	db.CreateTable()

	loopEvents(kafka, db)
	g.Run(":8080")
}

func loopEvents(kafka *internal.Pubsub, db *config.DBConn) {

	userRef := repository.NewUserRepository(db)

	i := 1
	for {

		userRef.Create(&models.User{
			Name: "leedonggyu",
			Age:  i,
		})

		kafka.Producer("home", map[string]any{
			"name":    "leedonggyu",
			"value":   i,
			"result2": i * i,
		})

		i++
		// time.Sleep(time.Second / 2)
		time.Sleep(time.Second * 5)
	}
}
