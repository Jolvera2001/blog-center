package main

import (
	"blog-center/internal/repository"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func test() {
	app := fx.New(
		fx.Provide(
			NewEnv,
			NewDB,
			NewRouter,
		),
	)

	app.Run()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("env file not found, going with environment variables")
	}

	port := os.Getenv("PORT")
	buildEnv := os.Getenv("GO_ENV")

	// connecting to DB
	maxRetries := 5
	maxDelay := 15 * time.Second
	_, err = repository.NewDB(maxRetries, maxDelay)
	if err != nil {
		log.Fatal("Error connecting to db!:", err.Error())
	}

	// setting up router
	if buildEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
		r := gin.New()
		r.Use(gin.Logger())

		r.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello World!",
			})
		})
		r.Run(port)
	} else {
		r := gin.Default()
		r.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello World!",
			})
		})
		r.Run(":" + port)
	}

}

