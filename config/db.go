package config

import (
	"backend-tugas-reactjs/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDataBase() *gorm.DB {
	var db *gorm.DB

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")

	dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + database + " port=" + port + " sslmode=disable"
	dbGorm, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db = dbGorm

	db.AutoMigrate(&models.Books{}, &models.User{}, &models.MataKuliah{}, &models.Mahasiswa{}, &models.Dosen{}, &models.JadwalKuliah{}, &models.Nilai{})

	return db

}
