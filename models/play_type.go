package models

import (
	"cfbapi/conn"
	"cfbapi/util"
	"log"
)

type PlayTypes []PlayType
type PlayType struct {
	Id           int    `json:"id"`
	Text         string `json:"text"`
	Abbreviation string `json:"abbreviation"`
}

func FetchAndInsertPlayTypes() error {
	var pt PlayTypes

	conn.APICall("plays/types", &pt)
	util.LogDBError("FetchAndInsertPlayTypes", conn.BatchInsert(util.CONFIG.CONNECTIONS.DB, pt, 1))
	log.Println("Inserted", len(pt), "play type records.")

	return nil
}
