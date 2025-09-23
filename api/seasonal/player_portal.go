package seasonal

import (
	"cfbapi/api"
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

func GetPortal() {
	b, _ := api.APICall("player/portal?year=" + strconv.Itoa(util.SEASON))
	var portal PortalSeason
	if err := json.Unmarshal(b, &portal); err != nil {
		panic(err)
	}

	InsertPortal()
}

func InsertPortal() {

}
