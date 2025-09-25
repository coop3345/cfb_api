package models

import (
	"cfbapi/conn"
	"cfbapi/util"
	"encoding/json"
)

type Conferences []Conference
type Conference struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	ShortName      string `json:"shortName"`
	Abbreviation   string `json:"abbreviation"`
	Classification string `json:"classification"`
}

func FetchAndInsertConferences() error {
	var con Conferences

	b, _ := conn.APICall("conferences")
	if err := json.Unmarshal(b, &con); err != nil {
		panic(err)
	}
	if err := util.DB.CreateInBatches(con, 100).Error; err != nil {
		return err
	}

	return nil
}
