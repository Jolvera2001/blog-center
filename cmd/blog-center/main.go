package main

import (
	"blog-center/internal/repository"
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type Env struct {
	Port     string
	BuildEnv string
}

func NewEnv() (*Env, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Env file not found, going with environment variables")
	}
	return &Env{
		Port: os.Getenv("PORT"),
		BuildEnv: os.Getenv("GO_ENV"),
	}, nil
}

func NewDB() (*gorm.DB, error) {
	maxRetries := 5
	maxDelay := 15 * time.Second
	return repository.NewDB(maxRetries, maxDelay)
}

func NewRouter(lc fx.Lifecycle, env *Env) *gin.Engine {
	var r *gin.Engine

	if env.BuildEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
		r.Use(gin.Logger())
	} else {
		r = gin.Default()
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				err := r.Run(":" + env.Port)
				if err != nil {
					log.Fatalf("Failed to start server: %v", err)
				}
			}()
			log.Println("Server started on port: ", env.Port)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Shutting down server...")
			return nil
		},
	})

	return r
}

func main() {
	app := fx.New(
		fx.Provide(
			NewEnv,
			NewDB,
			NewRouter,
		),
		fx.Invoke(func(*gin.Engine) {}),
	)

	app.Run()
}

// func main() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Println("env file not found, going with environment variables")
// 	}

// 	port := os.Getenv("PORT")
// 	buildEnv := os.Getenv("GO_ENV")

// 	// connecting to DB
// 	maxRetries := 5
// 	maxDelay := 15 * time.Second
// 	_, err = repository.NewDB(maxRetries, maxDelay)
// 	if err != nil {
// 		log.Fatal("Error connecting to db!:", err.Error())
// 	}

// 	// setting up router
// 	if buildEnv == "production" {
// 		gin.SetMode(gin.ReleaseMode)
// 		r := gin.New()
// 		r.Use(gin.Logger())

// 		r.GET("/", func(c *gin.Context) {
// 			c.JSON(http.StatusOK, gin.H{
// 				"message": "Hello World!",
// 			})
// 		})
// 		r.Run(port)
// 	} else {
// 		r := gin.Default()
// 		r.GET("/", func(c *gin.Context) {
// 			c.JSON(http.StatusOK, gin.H{
// 				"message": "Hello World!",
// 			})
// 		})
// 		r.Run(":" + port)
// 	}

// }
