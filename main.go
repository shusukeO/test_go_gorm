package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Project struct {
	gorm.Model
	Name     string
	RegionID string `gorm:"not null;default:'japan';"`
}

func main() {
	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Project{})

	db.Create(&Project{Name: "Test Project"})

	var project Project
	db.First(&project, 1)

	fmt.Println("Project ", project)
}

// connectDB はデータベースタイプに応じてDBに接続する
func connectDB() (*gorm.DB, error) {
	dbType := os.Getenv("DB_TYPE")
	if dbType == "" {
		dbType = "sqlite"
	}

	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	switch dbType {
	case "mysql":
		// MySQL接続設定
		dsn := "testuser:testpass@tcp(localhost:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), config)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to MySQL: %v", err)
		}
		fmt.Println("Connected to MySQL")
		return db, nil
	case "sqlite":
		// SQLite接続設定
		db, err := gorm.Open(sqlite.Open("test.db"), config)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to SQLite: %v", err)
		}
		fmt.Println("Connected to SQLite")
		return db, nil
	default:
		return nil, fmt.Errorf("unsupported DB_TYPE: %s", dbType)
	}
}
