package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	// dsn (data source name), basically the database connection string
	dsn := os.Getenv("DB_URL")
	// we are using gorm.Open to initialize a database session by using postgres.Open to use the postgres dialect
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err.Error())
	}
}
