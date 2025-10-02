package conn

import (
	"cfbapi/util"
	"errors"
	"fmt"
	"log"
	"reflect"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := util.DSN_STRING
	var err error
	DB, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}

	return DB, err
}

const maxSQLServerParams = 2100

func BatchInsert[T any](db *gorm.DB, data []T, batchSize int) error {
	if len(data) == 0 {
		return nil
	}

	var allErrors []error

	// Estimate number of fields per row
	var numFields int
	{
		// Use reflection to get field count from the first item
		sample := reflect.TypeOf(data[0])
		if sample.Kind() == reflect.Ptr {
			sample = sample.Elem()
		}
		numFields = sample.NumField()
		if numFields == 0 {
			return errors.New("no fields found in struct")
		}
	}

	// Calculate max batch size to stay within parameter limit
	maxBatch := maxSQLServerParams / numFields
	if batchSize > maxBatch {
		batchSize = maxBatch - 1 // keep under
	}

	for i := 0; i < len(data); i += batchSize {
		end := i + batchSize
		if end > len(data) {
			end = len(data)
		}

		batch := data[i:end]
		if err := db.CreateInBatches(batch, batchSize).Error; err != nil {
			allErrors = append(allErrors, fmt.Errorf("batch %d-%d failed: %w", i, end, err))
		}
	}

	if len(allErrors) > 0 {
		return fmt.Errorf("batch insert completed with errors: %v", allErrors)
	}

	return nil
}
