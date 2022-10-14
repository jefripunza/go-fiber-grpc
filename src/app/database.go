package app

import (
	"fmt"
	"os"

	"go-fiber-grpc/src/configs"
	"go-fiber-grpc/src/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
)

func DatabaseConnect() *gorm.DB {
	db_type := os.Getenv("DB_TYPE")
	fmt.Println(db_type)
	var connection *gorm.DB
	var err_msg error
	dsn := configs.DatabaseConfig()
	if db_type == "sqlite" {
		db, err := gorm.Open(sqlite.Open(fmt.Sprintf("%v.db", os.Getenv("DB_NAME"))), &gorm.Config{})
		connection = db
		err_msg = err
	} else if db_type == "mysql" {
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		connection = db
		err_msg = err
	} else if db_type == "postgres" {
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		connection = db
		err_msg = err
	}
	if err_msg != nil {
		log.Fatalln(err_msg)
		panic("failed to Database database")
	}
	return connection
}

func DatabaseMigration() {
	db_sync := os.Getenv("DB_SYNC")
	if db_sync == "true" {
		fmt.Println("Database Migrating...")
		db := DatabaseConnect()
		models.MigrationTables(db)
		fmt.Println("Migrate Success!")
	} else {
		fmt.Println("Database Off Migrating!")
	}
}
