package main

import (
	"fmt"
	"lark-gitlab-bridge/config"
	"lark-gitlab-bridge/entity"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func initDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%v&loc=%s",
		config.GlobalConfig.Database.Username,
		config.GlobalConfig.Database.Password,
		config.GlobalConfig.Database.Host,
		config.GlobalConfig.Database.Port,
		config.GlobalConfig.Database.Dbname,
		config.GlobalConfig.Database.Charset,
		config.GlobalConfig.Database.ParseTime,
		config.GlobalConfig.Database.Loc,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("failed to get database connection: %v", err)
	}
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	log.Println("Successfully connected to the database")
}

func runMigration() {
	migrateDSN := fmt.Sprintf("mysql://%s:%s@tcp(%s:%d)/%s",
		config.GlobalConfig.Database.Username,
		config.GlobalConfig.Database.Password,
		config.GlobalConfig.Database.Host,
		config.GlobalConfig.Database.Port,
		config.GlobalConfig.Database.Dbname,
	)

	m, err := migrate.New(
		"file://migrations",
		migrateDSN,
	)
	if err != nil {
		log.Fatalf("Error creating migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Error applying migration: %v", err)
	}

	log.Println("Migration applied successfully!")
}

func main() {
	initDB()
	runMigration()

	fmt.Println(config.GlobalConfig)
	fmt.Println("Hello, world!")

	name := "Test WebHook"
	url := "https://example.com/webhook"
	requiredVerification := false

	webhook := entity.Webhook{
		Name:                 name,
		URL:                  url,
		RequiredVerification: requiredVerification,
	}
	if err := webhook.ValidSecretKey(); err != nil {
		fmt.Println(err)
	} else {
		DB.Create(&webhook)
	}

}
