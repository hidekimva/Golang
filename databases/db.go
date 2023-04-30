package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaBD(host string, user string, password string, dbname string, port int) {
	stringDeConexao := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)

	DB, err = gorm.Open(postgres.Open(stringDeConexao))

	if err != nil {
		log.Panic("Failed connecting to database")
	}

	log.Println("Success connecting to database")

	DB.AutoMigrate(&User{})

}
