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
	// Define raw structure matching JSON input
	var raw struct {
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

	// Unmarshal into raw structure
	if err := json.Unmarshal(data, &raw); err != nil {
		return fmt.Errorf("failed to unmarshal GameTeamStat: %w", err)
	}

	var flat []GameTeamStat

	// Flatten loop
	for _, team := range raw.Teams {
		for _, stat := range team.Stats {
			flat = append(flat, GameTeamStat{
				GameID:     raw.Id,
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

	// Assign flattened results
	*gts = flat
	return nil
}

func FetchAndInsertGameTeamStats() error {
	var gts GameTeamStats
	query := fmt.Sprintf("games/teams?year=%v&week=%v", strconv.Itoa(util.SEASON), strconv.Itoa(util.WEEK))
	query = util.Trim_endpoint(query)

	b, _ := conn.APICall(query)
	if err := json.Unmarshal(b, &gts); err != nil {
		panic(err)
	}
	if err := util.DB.CreateInBatches(gts, 100).Error; err != nil {
		return err
	}

	return nil
}
