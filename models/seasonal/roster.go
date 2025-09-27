package seasonal

import (
	"cfbapi/conn"
	"cfbapi/util"
	"fmt"
	"strconv"

	"gorm.io/datatypes"
)

type Rosters []Roster
type Roster struct {
	Id             string         `json:"id" gorm:"primaryKey"`
	FirstName      string         `json:"firstName"`
	LastName       string         `json:"lastName"`
	Team           string         `json:"team" gorm:"primaryKey"`
	Season         int            `json:"season" gorm:"primaryKey"`
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

func FetchAndInsertRosters() error {
	var rosters Rosters
	query := fmt.Sprintf("roster?year=%v", strconv.Itoa(util.SEASON))
	conn.APICall(query, &rosters)
	for i := range rosters {
		rosters[i].Season = util.SEASON
	}

	if err := util.DB.CreateInBatches(rosters, 100).Error; err != nil {
		return err
	}

	return nil
}
