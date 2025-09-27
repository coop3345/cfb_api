package weekly

import (
	"cfbapi/conn"
	"cfbapi/util"
	"encoding/json"
	"fmt"
	"strconv"
)

// todo flatten
type GamePlayerStats []GamePlayerStat
type GamePlayerStat struct {
	GameID      int
	Team        string
	Conference  string
	HomeAway    string
	Points      int
	Category    string
	StatType    string
	AthleteID   string
	AthleteName string
	Stat        string
}

func (gps *GamePlayerStats) UnmarshalJSON(data []byte) error {
	// Define the raw JSON structure
	var raw struct {
		Id    int `json:"id"`
		Teams []struct {
			Team       string `json:"team"`
			Conference string `json:"conference"`
			HomeAway   string `json:"homeAway"`
			Points     int    `json:"points"`
			Categories []struct {
				Name  string `json:"name"`
				Types []struct {
					Name     string `json:"name"`
					Athletes []struct {
						Id   string `json:"id"`
						Name string `json:"name"`
						Stat string `json:"stat"`
					} `json:"athletes"`
				} `json:"types"`
			} `json:"categories"`
		} `json:"teams"`
	}

	// Unmarshal into raw form
	if err := json.Unmarshal(data, &raw); err != nil {
		return fmt.Errorf("failed to unmarshal GamePlayerStat: %w", err)
	}

	var flat []GamePlayerStat

	// Flatten the structure
	for _, team := range raw.Teams {
		for _, category := range team.Categories {
			for _, statType := range category.Types {
				for _, athlete := range statType.Athletes {
					flat = append(flat, GamePlayerStat{
						GameID:      raw.Id,
						Team:        team.Team,
						Conference:  team.Conference,
						HomeAway:    team.HomeAway,
						Points:      team.Points,
						Category:    category.Name,
						StatType:    statType.Name,
						AthleteID:   athlete.Id,
						AthleteName: athlete.Name,
						Stat:        athlete.Stat,
					})
				}
			}
		}
	}

	// Assign flattened result
	*gps = flat
	return nil
}

func FetchAndInsertGamePlayerStats() error {
	var gps GamePlayerStats
	query := fmt.Sprintf("games/players?year=%v&week=%v&seasonType=%v", strconv.Itoa(util.SEASON), strconv.Itoa(util.WEEK), util.SEASON_TYPE)
	query = util.Trim_endpoint(query)
	conn.APICall(query, &gps)
	if err := util.DB.CreateInBatches(gps, 100).Error; err != nil {
		return err
	}

	return nil
}
