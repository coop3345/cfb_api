package models

import (
	"cfbapi/conn"
	"cfbapi/util"
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
	conn.APICall("venues", &v)
	util.LogDBError("FetchAndInsertVenues", conn.BatchInsert(util.CONFIG.CONNECTIONS.DB, v, 1))

	return nil
}
