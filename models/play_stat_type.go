package models

import (
	"cfbapi/conn"
	"cfbapi/util"
	"encoding/json"
)

type PlayStatTypes []PlayStatType
type PlayStatType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func FetchAndInsertPlayStatTypes() error {
	var pst PlayStatTypes

	b, _ := conn.APICall("plays/stats/types")
	if err := json.Unmarshal(b, &pst); err != nil {
		panic(err)
	}
	if err := util.DB.CreateInBatches(pst, 100).Error; err != nil {
		return err
	}

	return nil
}
