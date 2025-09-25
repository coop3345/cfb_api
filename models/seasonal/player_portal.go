package seasonal

import (
	"cfbapi/conn"
	"cfbapi/util"
	"encoding/json"
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

func FetchAndInsertPortal() error {
	var portal PortalSeason
	query := "player/portal?year=" + strconv.Itoa(util.SEASON)

	b, _ := conn.APICall(query)
	if err := json.Unmarshal(b, &portal); err != nil {
		panic(err)
	}
	if err := util.DB.CreateInBatches(portal, 100).Error; err != nil {
		return err
	}

	return nil
}
