package seasonal

import (
	"cfbapi/conn"
	"cfbapi/util"
	"encoding/json"
	"fmt"
	"strconv"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Rosters []Roster
type Roster struct {
	Id             string         `json:"id" gorm:"primaryKey"`
	FirstName      string         `json:"firstName"`
	LastName       string         `json:"lastName"`
	Team           string         `json:"team"`
	Season         int            `json:"season"  gorm:"primaryKey"`
	Height         int            `json:"height"`
	Weight         int            `json:"weight"`
	Jersey         int            `json:"jersey"`
	Position       string         `json:"position"`
	HomeCity       string         `json:"homeCity"`
	HomeState      string         `json:"homeState"`
	HomeCountry    string         `json:"homeCountry"`
	HomeLatitude   float64        `json:"homeLatitude"`
	HomeLongitude  float64        `json:"homeLongitude"`
	HomeCountyFIPS string         `json:"homeCountyFIPS"`
	RecruitIds     datatypes.JSON `json:"recruitIds"`
}

func FetchAndInsertRosters(db *gorm.DB) error {
	var rosters Rosters
	query := fmt.Sprintf("roster?year=%v", strconv.Itoa(util.SEASON))
	b, _ := conn.APICall(query)
	if err := json.Unmarshal(b, &rosters); err != nil {
		return err
	}

	// Inject the season (from query context) into each item
	for i := range rosters {
		rosters[i].Season = util.SEASON
	}

	// Batch insert
	if err := db.CreateInBatches(rosters, 100).Error; err != nil {
		return err
	}

	return nil
}
