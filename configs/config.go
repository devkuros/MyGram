package configs

import (
	"mygram/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Conns() *gorm.DB {
	dsn := "host=localhost user=postgres password=root dbname=mygramdb port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.SocialMedia{})
	return db
}
