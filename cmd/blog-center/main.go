package main

import (
	"blog-center/internal/repository"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading dotenv!: ", err.Error())
	}

	maxRetries := 5
	maxDelay := 5 * time.Second
	_, err = repository.NewDB(maxRetries, maxDelay)
	if err != nil {
		log.Fatal("Error connecting to db!:", err.Error())
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})
	r.Run()
}

