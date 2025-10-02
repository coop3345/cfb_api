package weekly

import (
	"cfbapi/conn"
	"cfbapi/util"
	"encoding/json"
	"fmt"
	"strconv"
)

var GameEndpoint = "/games/teams"

type GameTeamStats []GameTeamStat
type GameTeamStat struct {
	GameID     int
	TeamID     int
	Team       string
	Conference string
	HomeAway   string
	Points     int
	Category   string
	Stat       string
}

func (gts *GameTeamStats) UnmarshalJSON(data []byte) error {
	// Define the raw input structure
	var rawGames []struct {
		Id    int `json:"id"`
		Teams []struct {
			TeamId     int    `json:"teamId"`
			Team       string `json:"team"`
			Conference string `json:"conference"`
			HomeAway   string `json:"homeAway"`
			Points     int    `json:"points"`
			Stats      []struct {
				Category string `json:"category"`
				Stat     string `json:"stat"`
			} `json:"stats"`
		} `json:"teams"`
	}

	// Try to unmarshal into a slice of raw game objects
	if err := json.Unmarshal(data, &rawGames); err != nil {
		return fmt.Errorf("failed to unmarshal GameTeamStat: %w", err)
	}

	var flat []GameTeamStat

	for _, game := range rawGames {
		for _, team := range game.Teams {
			for _, stat := range team.Stats {
				flat = append(flat, GameTeamStat{
					GameID:     game.Id,
					TeamID:     team.TeamId,
					Team:       team.Team,
					Conference: team.Conference,
					HomeAway:   team.HomeAway,
					Points:     team.Points,
					Category:   stat.Category,
					Stat:       stat.Stat,
				})
			}
		}
	}

	*gts = flat
	return nil
}

func FetchAndInsertGameTeamStats() error {
	var gts GameTeamStats
	query := fmt.Sprintf("games/teams?year=%v&week=%v&seasonType=%v", strconv.Itoa(util.SEASON), strconv.Itoa(util.WEEK), util.SEASON_TYPE)
	query = util.Trim_endpoint(query)
	conn.APICall(query, &gts)
	util.LogDBError("FetchAndInsertGameTeamStats", conn.BatchInsert(util.DB, gts, 100))
	return nil
}
