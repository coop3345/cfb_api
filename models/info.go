package models

import (
	"cfbapi/conn"
	"cfbapi/util"
	"log"
	"time"
)

type Info struct {
	PatronLevel    int `json:"patronLevel"`
	RemainingCalls int `json:"remainingCalls"`
	Timestamp      time.Time
}

func (Info) TableName() string {
	return "api_info"
}

func FetchAndInsertInfo() error {
	var info Info
	query := "info"
	err := conn.APICall(query, &info)
	if err != nil {
		return err
	}

	info.Timestamp = time.Now()

	infoSlice := []Info{info}
	if err := conn.BatchInsert(util.CONFIG.CONNECTIONS.DB, infoSlice, 100); err != nil {
		return err
	}

	log.Println("Inserted info record.")

	return nil
}
