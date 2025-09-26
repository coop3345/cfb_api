package seasonal

import (
	"cfbapi/conn"
	"cfbapi/util"
	"encoding/json"
	"fmt"
	"strconv"

	"gorm.io/datatypes"
)

var TEAMS Teams
var CONFERENCE_TEAMS map[string]Teams

type Teams []Team
type Team struct {
	ID             int            `json:"id" gorm:"primaryKey"`
	School         string         `json:"school"`
	Mascot         string         `json:"mascot"`
	Abbreviation   string         `json:"abbreviation"`
	AlternateNames datatypes.JSON `json:"alternateNames"`
	Conference     string         `json:"conference"`
	Division       string         `json:"division"`
	Classification string         `json:"classification"`
	Color          string         `json:"color"`
	AlternateColor string         `json:"alternateColor"`
	Logos          datatypes.JSON `json:"logos"`
	Twitter        string         `json:"twitter"`

	// Flattened Location fields
	LocationID       int     `json:"-"`
	LocationName     string  `json:"-"`
	City             string  `json:"-"`
	State            string  `json:"-"`
	Zip              string  `json:"-"`
	CountryCode      string  `json:"-"`
	Timezone         string  `json:"-"`
	Latitude         float64 `json:"-"`
	Longitude        float64 `json:"-"`
	Elevation        string  `json:"-"`
	Capacity         int     `json:"-"`
	ConstructionYear int     `json:"-"`
	Grass            bool    `json:"-"`
	Dome             bool    `json:"-"`
}
type Location struct {
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

func (t *Team) UnmarshalJSON(data []byte) error {
	// Define alias to match original JSON structure
	type Alias Team
	aux := &struct {
		Location struct {
			ID               int     `json:"id"`
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
		} `json:"location"`
		*Alias
	}{
		Alias: (*Alias)(t), // embed all other fields
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return fmt.Errorf("failed to unmarshal Team: %w", err)
	}

	// Flatten location into main struct
	t.LocationID = aux.Location.ID
	t.LocationName = aux.Location.Name
	t.City = aux.Location.City
	t.State = aux.Location.State
	t.Zip = aux.Location.Zip
	t.CountryCode = aux.Location.CountryCode
	t.Timezone = aux.Location.Timezone
	t.Latitude = aux.Location.Latitude
	t.Longitude = aux.Location.Longitude
	t.Elevation = aux.Location.Elevation
	t.Capacity = aux.Location.Capacity
	t.ConstructionYear = aux.Location.ConstructionYear
	t.Grass = aux.Location.Grass
	t.Dome = aux.Location.Dome

	return nil
}

func FetchAndInsertTeams() error {
	query := fmt.Sprintf("teams?year=%v", strconv.Itoa(util.SEASON))
	conn.APICall(query, &TEAMS)
	if err := util.DB.CreateInBatches(TEAMS, 100).Error; err != nil {
		return err
	}

	BuildConferenceTeams()

	return nil
}

func BuildConferenceTeams() {
	CONFERENCE_TEAMS = make(map[string]Teams)

	for _, team := range TEAMS {
		conf := team.Conference
		CONFERENCE_TEAMS[conf] = append(CONFERENCE_TEAMS[conf], team)
	}
}
