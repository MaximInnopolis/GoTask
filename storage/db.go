package storage

import (
	"TestTask/models"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var DB *gorm.DB

// load .env file
func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func DBConn() *gorm.DB {
	// pass the DB credentials into variables
	host := goDotEnvVariable("DBHOST")
	port := goDotEnvVariable("DBPORT")
	dbUser := goDotEnvVariable("DBUSER")
	dbPass := goDotEnvVariable("DBPASS")
	dbname := goDotEnvVariable("DBNAME")
	// create a connection string
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, dbUser, dbPass, dbname)
	var err error

	DB, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	err = DB.AutoMigrate(&models.Check{}, &models.CheckHash{})
	if err != nil {
		log.Fatal(err)
	}

	return DB
}

func GetDBInstance() *gorm.DB {
	return DB
}
