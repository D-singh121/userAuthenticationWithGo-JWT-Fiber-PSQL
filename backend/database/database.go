package database

import (
	"fmt"

	"github.com/devesh/golang-react-jwt/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// dsn := "postgresql://admin:admin@localhost:5432/go"
	dsn := `host=localhost
			user=admin
			password=admin
			dbname=gofiberjwt
			port=5432
			sslmode=disable`
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database: " + err.Error())
	} else {
		fmt.Println("package database connected successfully!")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
