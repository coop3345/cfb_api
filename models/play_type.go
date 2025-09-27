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
	util.LogDBError("FetchAndInsertPlayTypes", util.DB.CreateInBatches(pt, 1).Error)

	return nil
}
