package seasonal

import (
	"cfbapi/conn"
	"cfbapi/util"
	"log"
	"strconv"
)

type PortalSeason []PlayerPortalEntry
type PlayerPortalEntry struct {
	Season       int     `json:"season"`
	FirstName    string  `json:"firstName"`
	LastName     string  `json:"lastName"`
	Position     string  `json:"position"`
	Origin       string  `json:"origin"`
	Destination  string  `json:"destination"`
	TransferDate string  `json:"transferDate"`
	Rating       float64 `json:"rating"`
	Stars        int     `json:"stars"`
	Eligibility  string  `json:"eligibility"`
}

func (PlayerPortalEntry) TableName() string {
	return "Portal"
}

func FetchAndInsertPortal() error {
	var portal PortalSeason
	query := "player/portal?year=" + strconv.Itoa(util.SEASON)
	conn.APICall(query, &portal)
	util.LogDBError("FetchAndInsertPortal", conn.BatchInsert(util.CONFIG.CONNECTIONS.DB, portal, 100))
	log.Println("Inserted", len(portal), "portal records.")

	return nil
}
