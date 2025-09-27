package models

import (
	"cfbapi/conn"
	"cfbapi/util"
)

var CONFERENCES Conferences

type Conferences []Conference
type Conference struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	ShortName      string `json:"shortName"`
	Abbreviation   string `json:"abbreviation"`
	Classification string `json:"classification"`
}

func FetchAndInsertConferences() error {
	conn.APICall("conferences", &CONFERENCES)
	util.LogDBError("FetchAndInsertConferences", util.DB.CreateInBatches(CONFERENCES, 1).Error)

	return nil
}
