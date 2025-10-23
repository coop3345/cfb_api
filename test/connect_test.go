package test

import (
	"cfbapi/conn"
	"cfbapi/models"
	"cfbapi/util"
	"log"
	"testing"
)

func TestConnect(t *testing.T) {
	var err error
	t.Run("setup config", func(t *testing.T) {
		err = util.Setup()
		if err != nil {
			t.Errorf("Failed to setup config: %v", err)
		}
	})

	t.Run("connect to database", func(t *testing.T) {
		util.CONFIG.CONNECTIONS.DB, err = conn.InitDB()
		if err != nil {
			t.Errorf("Failed to connect to database: %v", err)
		}
	})
	t.Run("migrate models", func(t *testing.T) {
		err = models.Migrate_Model()
		if err != nil {
			t.Errorf("Failed to migrate models: %v", err)
		}
	})
	t.Run("fetch and insert info", func(t *testing.T) {
		err = models.FetchAndInsertInfo()
		if err != nil {
			t.Errorf("Failed to fetch and insert info: %v", err)
		}
	})

	log.Println("Test completed successfully")
}
