package weekly

import (
	"cfbapi/conn"
	"cfbapi/util"
	"encoding/json"
	"fmt"
	"strconv"
)

type Drives []Drive
type Drive struct {
	Offense           string `json:"offense"`
	OffenseConference string `json:"offenseConference"`
	Defense           string `json:"defense"`
	DefenseConference string `json:"defenseConference"`
	GameId            int    `json:"gameId" gorm:"primaryKey"`
	Id                string `json:"id" gorm:"primaryKey"`
	DriveNumber       int    `json:"driveNumber"`
	Scoring           bool   `json:"scoring"`
	StartPeriod       int    `json:"startPeriod"`
	StartYardline     int    `json:"startYardline"`
	StartYardsToGoal  int    `json:"startYardsToGoal"`
	StartMinutes      int
	StartSeconds      int
	EndPeriod         int `json:"endPeriod"`
	EndYardline       int `json:"endYardline"`
	EndYardsToGoal    int `json:"endYardsToGoal"`
	EndMinutes        int
	EndSeconds        int
	ElapsedMinutes    int
	ElapsedSeconds    int
	Plays             int    `json:"plays"`
	Yards             int    `json:"yards"`
	DriveResult       string `json:"driveResult"`
	IsHomeOffense     bool   `json:"isHomeOffense"`
	StartOffenseScore int    `json:"startOffenseScore"`
	StartDefenseScore int    `json:"startDefenseScore"`
	EndOffenseScore   int    `json:"endOffenseScore"`
	EndDefenseScore   int    `json:"endDefenseScore"`
}

func (d *Drive) UnmarshalJSON(data []byte) error {
	type timeObj struct {
		Minutes int `json:"minutes"`
		Seconds int `json:"seconds"`
	}

	// Alias the Drive and override time fields
	type Alias Drive
	aux := &struct {
		StartTime timeObj `json:"startTime"`
		EndTime   timeObj `json:"endTime"`
		Elapsed   timeObj `json:"elapsed"`
		*Alias
	}{
		Alias: (*Alias)(d),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return fmt.Errorf("failed to unmarshal Drive: %w", err)
	}

	// Flatten time fields
	d.StartMinutes = aux.StartTime.Minutes
	d.StartSeconds = aux.StartTime.Seconds
	d.EndMinutes = aux.EndTime.Minutes
	d.EndSeconds = aux.EndTime.Seconds
	d.ElapsedMinutes = aux.Elapsed.Minutes
	d.ElapsedSeconds = aux.Elapsed.Seconds

	return nil
}

func FetchAndInsertDrives() error {
	var drives Drives
	query := fmt.Sprintf("drives?year=%v&week=%v", strconv.Itoa(util.SEASON), strconv.Itoa(util.WEEK))
	query = util.Trim_endpoint(query)
	conn.APICall(query, &drives)
	if err := util.DB.CreateInBatches(drives, 100).Error; err != nil {
		return err
	}

	return nil
}
