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
	// データベースタイプを環境変数で切り替え (DB_TYPE=mysql または DB_TYPE=sqlite)
	dbType := os.Getenv("DB_TYPE")
	if dbType == "" {
		dbType = "sqlite"
	}

	var db *gorm.DB
	var err error

	switch dbType {
	case "mysql":
		// MySQL接続設定
		dsn := "testuser:testpass@tcp(localhost:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			log.Fatalf("failed to connect to MySQL: %v", err)
		}
		fmt.Println("Connected to MySQL")
	case "sqlite":
		// SQLite接続設定
		db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			log.Fatalf("failed to connect to SQLite: %v", err)
		}
		fmt.Println("Connected to SQLite")
	default:
		log.Fatalf("Unsupported DB_TYPE: %s", dbType)
	}

	db.AutoMigrate(&Project{})

	// db.Create(&Project{Name: "Test Project" /*RegionID: "japan"*/})

	var project Project
	db.First(&project, 1)

	fmt.Println("Project ", project)
}
