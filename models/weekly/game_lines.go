package weekly

import (
	"cfbapi/conn"
	"cfbapi/util"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type GameLines []FlatGameLine
type FlatGameLine struct {
	GameID             int       `json:"game_id"`
	Season             int       `json:"season"`
	SeasonType         string    `json:"season_type"`
	Week               int       `json:"week"`
	StartDate          time.Time `json:"start_date"`
	HomeTeamID         int       `json:"home_team_id"`
	HomeTeam           string    `json:"home_team"`
	HomeConference     string    `json:"home_conference"`
	HomeClassification string    `json:"home_classification"`
	HomeScore          int       `json:"home_score"`
	AwayTeamID         int       `json:"away_team_id"`
	AwayTeam           string    `json:"away_team"`
	AwayConference     string    `json:"away_conference"`
	AwayClassification string    `json:"away_classification"`
	AwayScore          int       `json:"away_score"`
	Provider           string    `json:"provider"`
	Spread             float64   `json:"spread"`
	FormattedSpread    string    `json:"formatted_spread"`
	SpreadOpen         float64   `json:"spread_open"`
	OverUnder          float64   `json:"over_under"`
	OverUnderOpen      float64   `json:"over_under_open"`
	HomeMoneyline      int       `json:"home_moneyline"`
	AwayMoneyline      int       `json:"away_moneyline"`
}

func (FlatGameLine) TableName() string {
	return "game_lines"
}

func (gts *GameLines) UnmarshalJSON(data []byte) error {
	// Define the raw nested structure
	type rawGameLines struct {
		ID                 int    `json:"id"`
		Season             int    `json:"season"`
		SeasonType         string `json:"seasonType"`
		Week               int    `json:"week"`
		StartDate          string `json:"startDate"`
		HomeTeamID         int    `json:"homeTeamId"`
		HomeTeam           string `json:"homeTeam"`
		HomeConference     string `json:"homeConference"`
		HomeClassification string `json:"homeClassification"`
		HomeScore          int    `json:"homeScore"`
		AwayTeamID         int    `json:"awayTeamId"`
		AwayTeam           string `json:"awayTeam"`
		AwayConference     string `json:"awayConference"`
		AwayClassification string `json:"awayClassification"`
		AwayScore          int    `json:"awayScore"`
		Lines              []struct {
			Provider        string  `json:"provider"`
			Spread          float64 `json:"spread"`
			FormattedSpread string  `json:"formattedSpread"`
			SpreadOpen      float64 `json:"spreadOpen"`
			OverUnder       float64 `json:"overUnder"`
			OverUnderOpen   float64 `json:"overUnderOpen"`
			HomeMoneyline   int     `json:"homeMoneyline"`
			AwayMoneyline   int     `json:"awayMoneyline"`
		} `json:"lines"`
	}

	var rawGames []rawGameLines
	if err := json.Unmarshal(data, &rawGames); err != nil {
		return err
	}

	var flat GameLines
	for _, g := range rawGames {
		startTime, err := time.Parse(time.RFC3339, g.StartDate)
		if err != nil {
			return fmt.Errorf("invalid start date: %w", err)
		}

		for _, line := range g.Lines {
			flat = append(flat, FlatGameLine{
				GameID:             g.ID,
				Season:             g.Season,
				SeasonType:         g.SeasonType,
				Week:               g.Week,
				StartDate:          startTime,
				HomeTeamID:         g.HomeTeamID,
				HomeTeam:           g.HomeTeam,
				HomeConference:     g.HomeConference,
				HomeClassification: g.HomeClassification,
				HomeScore:          g.HomeScore,
				AwayTeamID:         g.AwayTeamID,
				AwayTeam:           g.AwayTeam,
				AwayConference:     g.AwayConference,
				AwayClassification: g.AwayClassification,
				AwayScore:          g.AwayScore,
				Provider:           line.Provider,
				Spread:             line.Spread,
				FormattedSpread:    line.FormattedSpread,
				SpreadOpen:         line.SpreadOpen,
				OverUnder:          line.OverUnder,
				OverUnderOpen:      line.OverUnderOpen,
				HomeMoneyline:      line.HomeMoneyline,
				AwayMoneyline:      line.AwayMoneyline,
			})
		}
	}

	*gts = flat
	return nil
}

func FetchAndInsertGameLines() error {
	var game_lines GameLines
	query := fmt.Sprintf("lines?year=%v&week=%v&seasonType=%v", strconv.Itoa(util.SEASON), strconv.Itoa(util.WEEK), util.SEASON_TYPE)
	query = util.Trim_endpoint(query)
	conn.APICall(query, &game_lines)
	util.LogDBError("FetchAndInsertDrives", conn.BatchInsert(util.CONFIG.CONNECTIONS.DB, game_lines, 200))

	return nil
}

func FetchAndInsertGameLinesYear() error {
	var game_lines GameLines
	query := fmt.Sprintf("lines?year=%v", strconv.Itoa(util.SEASON))
	conn.APICall(query, &game_lines)
	util.LogDBError("FetchAndInsertDrives", conn.BatchInsert(util.CONFIG.CONNECTIONS.DB, game_lines, 200))

	return nil
}
