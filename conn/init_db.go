package conn

import (
	"cfbapi/models"
	"cfbapi/util"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := util.DSN_STRING
	var err error
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}

	DB.AutoMigrate(&models.Recruit{})

	return DB, err
}
