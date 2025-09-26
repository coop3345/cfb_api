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
	if err := util.DB.CreateInBatches(pst, 100).Error; err != nil {
		return err
	}

	return nil
}
