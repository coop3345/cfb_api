package models

import (
	"cfbapi/conn"
	"cfbapi/util"
)

var CONFERENCES Conferences

// var COLLECT_CONFERENCES []string

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
	util.LogDBError("FetchAndInsertConferences", conn.BatchInsert(util.DB, CONFERENCES, 1))

	// for _, con := range CONFERENCES {
	// 	if util.Contains(util.PSCD, con.Classification) {
	// 		COLLECT_CONFERENCES = append(COLLECT_CONFERENCES, con.Name)
	// 	}
	// }

	return nil
}
