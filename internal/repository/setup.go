package repository

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(retries int, retryDelay time.Duration) (*gorm.DB, error) {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		return nil, fmt.Errorf("DB_DSN environment variable not set!")
	}

	for i := 0; i < retries; i++ {
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Printf("Failed to connect to database: %v, Retrying in %v...", err, retryDelay)
			time.Sleep(retryDelay)
			continue
		}

		sqlDB, err := db.DB()
		if err != nil {
			return nil, err
		}	
	
		err = sqlDB.Ping()
		if err == nil {
			fmt.Println("Successfully connected to database")
			return db, nil
		}

		fmt.Printf("Failed to ping database: %v. Retrying in %v...\n", err, retryDelay)
        time.Sleep(retryDelay)
	}


	return nil, fmt.Errorf("failed to connect to the database after %d attempts", retries)
}
