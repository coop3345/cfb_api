package models

import (
	"cfbapi/conn"
	"cfbapi/util"
	"encoding/json"
)

type Venues []Venue
type Venue struct {
	Id               int     `json:"id"`
	Name             string  `json:"name"`
	City             string  `json:"city"`
	State            string  `json:"state"`
	Zip              string  `json:"zip"`
	CountryCode      string  `json:"countryCode"`
	Timezone         string  `json:"timezone"`
	Latitude         float64 `json:"latitude"`
	Longitude        float64 `json:"longitude"`
	Elevation        string  `json:"elevation"`
	Capacity         int     `json:"capacity"`
	ConstructionYear int     `json:"constructionYear"`
	Grass            bool    `json:"grass"`
	Dome             bool    `json:"dome"`
}

func FetchAndInsertVenues() error {
	var v Venues

	b, _ := conn.APICall("venues")
	if err := json.Unmarshal(b, &v); err != nil {
		panic(err)
	}
	if err := util.DB.CreateInBatches(v, 100).Error; err != nil {
		return err
	}

	return nil
}
