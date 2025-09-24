package weekly

import (
	"encoding/json"
	"fmt"
)

type Plays []Play
type Play struct {
	Id                string `json:"id"`
	DriveId           string `json:"driveId"`
	GameId            int    `json:"gameId"`
	DriveNumber       int    `json:"driveNumber"`
	PlayNumber        int    `json:"playNumber"`
	Offense           string `json:"offense"`
	OffenseConference string `json:"offenseConference"`
	OffenseScore      int    `json:"offenseScore"`
	Defense           string `json:"defense"`
	Home              string `json:"home"`
	Away              string `json:"away"`
	DefenseConference string `json:"defenseConference"`
	DefenseScore      int    `json:"defenseScore"`
	Period            int    `json:"period"`
	ClockMinutes      int
	ClockSeconds      int
	OffenseTimeouts   int     `json:"offenseTimeouts"`
	DefenseTimeouts   int     `json:"defenseTimeouts"`
	Yardline          int     `json:"yardline"`
	YardsToGoal       int     `json:"yardsToGoal"`
	Down              int     `json:"down"`
	Distance          int     `json:"distance"`
	YardsGained       int     `json:"yardsGained"`
	Scoring           bool    `json:"scoring"`
	PlayType          string  `json:"playType"`
	PlayText          string  `json:"playText"`
	Ppa               float64 `json:"ppa"`
	Wallclock         string  `json:"wallclock"`
}

func (p *Play) UnmarshalJSON(data []byte) error {
	type clockObj struct {
		Minutes int `json:"minutes"`
		Seconds int `json:"seconds"`
	}
	type Alias Play
	aux := &struct {
		Clock clockObj `json:"clock"`
		*Alias
	}{
		Alias: (*Alias)(p),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return fmt.Errorf("failed to unmarshal Play: %w", err)
	}

	p.ClockMinutes = aux.Clock.Minutes
	p.ClockSeconds = aux.Clock.Seconds

	return nil
}

type PlayStat struct {
	GameId        int    `json:"gameId"`
	Season        int    `json:"season"`
	Week          int    `json:"week"`
	Team          string `json:"team"`
	Conference    string `json:"conference"`
	Opponent      string `json:"opponent"`
	TeamScore     int    `json:"teamScore"`
	OpponentScore int    `json:"opponentScore"`
	DriveId       string `json:"driveId"`
	PlayId        string `json:"playId"`
	Period        int    `json:"period"`
	ClockMinutes  int
	ClockSeconds  int
	YardsToGoal   int    `json:"yardsToGoal"`
	Down          int    `json:"down"`
	Distance      int    `json:"distance"`
	AthleteId     string `json:"athleteId"`
	AthleteName   string `json:"athleteName"`
	StatType      string `json:"statType"`
	Stat          int    `json:"stat"`
}

func (ps *PlayStat) UnmarshalJSON(data []byte) error {
	type clockObj struct {
		Minutes int `json:"minutes"`
		Seconds int `json:"seconds"`
	}
	type Alias PlayStat
	aux := &struct {
		Clock clockObj `json:"clock"`
		*Alias
	}{
		Alias: (*Alias)(ps),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return fmt.Errorf("failed to unmarshal PlayStat: %w", err)
	}

	ps.ClockMinutes = aux.Clock.Minutes
	ps.ClockSeconds = aux.Clock.Seconds

	return nil
}
