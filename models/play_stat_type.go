package models

import (
	"cfbapi/conn"
	"cfbapi/util"
)

type PlayStatTypes []PlayStatType
type PlayStatType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func FetchAndInsertPlayStatTypes() error {
	var pst PlayStatTypes
	conn.APICall("plays/stats/types", &pst)
	util.LogDBError("FetchAndInsertPlayStatTypes", util.DB.CreateInBatches(pst, 1).Error)

	return nil
}
