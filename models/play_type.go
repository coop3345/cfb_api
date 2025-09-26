package models

import (
	"cfbapi/conn"
	"cfbapi/util"
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
	if err := util.DB.CreateInBatches(pt, 100).Error; err != nil {
		return err
	}

	return nil
}
