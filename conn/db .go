package conn

import (
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

	return DB, err
}

func BatchInsert[T any](db *gorm.DB, data []T, batchSize int) error {
	if len(data) == 0 {
		return nil
	}
	return db.CreateInBatches(data, batchSize).Error
}
