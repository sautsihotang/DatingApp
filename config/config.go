package config

import (
	"fmt"
	"log"

	"github.com/sautsihotang/DatingApp/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DBHost     = "localhost"
	DBPort     = 5432
	DBUser     = "postgres"
	DBPassword = "postgres"
	DBName     = "DatingAppDB"
)

var (
	DB *gorm.DB
)

func Init() {
	//membuat koneksi string
	dbURI := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		DBHost,
		DBPort,
		DBUser,
		DBPassword,
		DBName)

	//koneksi ke db
	var err error
	DB, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Berhasil terhubung ke database %s \n", DBName)

	//melakukan migrasi ke database
	MigrateModels()
}

func MigrateModels() {
	err := DB.AutoMigrate(&models.User{}, &models.Profile{}, &models.Swipe{})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Berhasil migrate ke database %s \n", DBName)
}
