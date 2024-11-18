package main

import (
	"context"
	"fmt"
	"log"

	"go-hire/config"
	"go-hire/ent"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/go-sql-driver/mysql"
)

var client *ent.Client

func initDB() {
	// MySQL에 연결
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
	client, err = ent.Open(dialect.MySQL, dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer client.Close()

	// 데이터베이스 스키마 마이그레이션
	if err := client.Schema.Create(context.Background(),
		schema.WithDropColumn(true),
		schema.WithDropIndex(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	log.Println("Migration applied successfully!")
}

func main() {
	// DB 초기화 및 마이그레이션
	initDB()

	fmt.Println("Hello, world!")
}
