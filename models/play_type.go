package models

import (
	"cfbapi/conn"
	"cfbapi/util"
	"encoding/json"
)

type PlayTypes []PlayType
type PlayType struct {
	Id           int    `json:"id"`
	Text         string `json:"text"`
	Abbreviation string `json:"abbreviation"`
}

func FetchAndInsertPlayTypes() error {
	var pt PlayTypes

	b, _ := conn.APICall("plays/types")
	if err := json.Unmarshal(b, &pt); err != nil {
		panic(err)
	}
	if err := util.DB.CreateInBatches(pt, 100).Error; err != nil {
		return err
	}

	return nil
}
