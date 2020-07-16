package models

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func ConnectDB() {
	connectPG(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"),
	)
	debug, _ := strconv.ParseBool(os.Getenv("DB_DEBUG"))
	db.LogMode(debug)
	fmt.Println("Debug mode:", debug)
}

func connectPG(DBDriver, DBUser, DBPassword, DBPort, DBHost, DBName string) {
	database, err := gorm.Open(DBDriver, fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		DBHost, DBPort, DBUser, DBName, DBPassword,
	))
	if err != nil {
		fmt.Println("Postgres cannot connect to", DBName)
		log.Fatal("Error", err)
	}
	fmt.Println("Postgres connect to", DBName)
	db = database
}
