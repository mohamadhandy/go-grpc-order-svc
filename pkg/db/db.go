package db

import (
	"go-grpc-order-svc/pkg/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect db: %v", err)
	}
	db.AutoMigrate(&models.Order{})
	return Handler{db}
}
