package database

import (
	"fmt"

	"github.com/dylanwe/yifu/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	c := config.GetConfig()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
		c.DBHost,
		c.DBUser,
		c.DBPassword,
		c.DBName,
	)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Clothing{})
	DB = database
}

func GetDB() *gorm.DB {
	return DB
}
