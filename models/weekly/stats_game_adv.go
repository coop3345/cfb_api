package weekly

import (
	"cfbapi/conn"
	"cfbapi/util"
	"encoding/json"
	"fmt"
	"strconv"
)

type StatsGameAdv []StatsGameAdvFlat
type StatsGameAdvFlat struct {
	GameId   int    `json:"gameId"`
	Season   int    `json:"season"`
	Week     int    `json:"week"`
	Team     string `json:"team"`
	Opponent string `json:"opponent"`

	// Offense PassingPlays
	OffensePassingPlaysExplosiveness float64 `json:"offensePassingPlaysExplosiveness"`
	OffensePassingPlaysSuccessRate   float64 `json:"offensePassingPlaysSuccessRate"`
	OffensePassingPlaysTotalPpa      float64 `json:"offensePassingPlaysTotalPPA"`
	OffensePassingPlaysPpa           float64 `json:"offensePassingPlaysPpa"`

	// Offense RushingPlays
	OffenseRushingPlaysExplosiveness float64 `json:"offenseRushingPlaysExplosiveness"`
	OffenseRushingPlaysSuccessRate   float64 `json:"offenseRushingPlaysSuccessRate"`
	OffenseRushingPlaysTotalPpa      float64 `json:"offenseRushingPlaysTotalPPA"`
	OffenseRushingPlaysPpa           float64 `json:"offenseRushingPlaysPpa"`

	// Offense PassingDowns
	OffensePassingDownsExplosiveness float64 `json:"offensePassingDownsExplosiveness"`
	OffensePassingDownsSuccessRate   float64 `json:"offensePassingDownsSuccessRate"`
	OffensePassingDownsPpa           float64 `json:"offensePassingDownsPpa"`

	// Offense StandardDowns
	OffenseStandardDownsExplosiveness float64 `json:"offenseStandardDownsExplosiveness"`
	OffenseStandardDownsSuccessRate   float64 `json:"offenseStandardDownsSuccessRate"`
	OffenseStandardDownsPpa           float64 `json:"offenseStandardDownsPpa"`

	// Offense Other Metrics
	OffenseOpenFieldYardsTotal   float64 `json:"offenseOpenFieldYardsTotal"`
	OffenseOpenFieldYards        float64 `json:"offenseOpenFieldYards"`
	OffenseSecondLevelYardsTotal float64 `json:"offenseSecondLevelYardsTotal"`
	OffenseSecondLevelYards      float64 `json:"offenseSecondLevelYards"`
	OffenseLineYardsTotal        float64 `json:"offenseLineYardsTotal"`
	OffenseLineYards             float64 `json:"offenseLineYards"`
	OffenseStuffRate             float64 `json:"offenseStuffRate"`
	OffensePowerSuccess          float64 `json:"offensePowerSuccess"`
	OffenseExplosiveness         float64 `json:"offenseExplosiveness"`
	OffenseSuccessRate           float64 `json:"offenseSuccessRate"`
	OffenseTotalPpa              float64 `json:"offenseTotalPpa"`
	OffensePpa                   float64 `json:"offensePpa"`
	OffenseDrives                int     `json:"offenseDrives"`
	OffensePlays                 int     `json:"offensePlays"`

	// Defense PassingPlays
	DefensePassingPlaysExplosiveness float64 `json:"defensePassingPlaysExplosiveness"`
	DefensePassingPlaysSuccessRate   float64 `json:"defensePassingPlaysSuccessRate"`
	DefensePassingPlaysTotalPpa      float64 `json:"defensePassingPlaysTotalPPA"`
	DefensePassingPlaysPpa           float64 `json:"defensePassingPlaysPpa"`

	// Defense RushingPlays
	DefenseRushingPlaysExplosiveness float64 `json:"defenseRushingPlaysExplosiveness"`
	DefenseRushingPlaysSuccessRate   float64 `json:"defenseRushingPlaysSuccessRate"`
	DefenseRushingPlaysTotalPpa      float64 `json:"defenseRushingPlaysTotalPPA"`
	DefenseRushingPlaysPpa           float64 `json:"defenseRushingPlaysPpa"`

	// Defense PassingDowns
	DefensePassingDownsExplosiveness float64 `json:"defensePassingDownsExplosiveness"`
	DefensePassingDownsSuccessRate   float64 `json:"defensePassingDownsSuccessRate"`
	DefensePassingDownsPpa           float64 `json:"defensePassingDownsPpa"`

	// Defense StandardDowns
	DefenseStandardDownsExplosiveness float64 `json:"defenseStandardDownsExplosiveness"`
	DefenseStandardDownsSuccessRate   float64 `json:"defenseStandardDownsSuccessRate"`
	DefenseStandardDownsPpa           float64 `json:"defenseStandardDownsPpa"`

	// Defense Other Metrics
	DefenseOpenFieldYardsTotal   float64 `json:"defenseOpenFieldYardsTotal"`
	DefenseOpenFieldYards        float64 `json:"defenseOpenFieldYards"`
	DefenseSecondLevelYardsTotal float64 `json:"defenseSecondLevelYardsTotal"`
	DefenseSecondLevelYards      float64 `json:"defenseSecondLevelYards"`
	DefenseLineYardsTotal        float64 `json:"defenseLineYardsTotal"`
	DefenseLineYards             float64 `json:"defenseLineYards"`
	DefenseStuffRate             float64 `json:"defenseStuffRate"`
	DefensePowerSuccess          float64 `json:"defensePowerSuccess"`
	DefenseExplosiveness         float64 `json:"defenseExplosiveness"`
	DefenseSuccessRate           float64 `json:"defenseSuccessRate"`
	DefenseTotalPpa              float64 `json:"defenseTotalPpa"`
	DefensePpa                   float64 `json:"defensePpa"`
	DefenseDrives                int     `json:"defenseDrives"`
	DefensePlays                 int     `json:"defensePlays"`
}

type StatsGameAdvRaw struct {
	GameId   int    `json:"gameId"`
	Season   int    `json:"season"`
	Week     int    `json:"week"`
	Team     string `json:"team"`
	Opponent string `json:"opponent"`
	Offense  struct {
		PassingPlays struct {
			Explosiveness float64 `json:"explosiveness"`
			SuccessRate   float64 `json:"successRate"`
			TotalPpa      float64 `json:"totalPPA"`
			Ppa           float64 `json:"ppa"`
		} `json:"passingPlays"`
		RushingPlays struct {
			Explosiveness float64 `json:"explosiveness"`
			SuccessRate   float64 `json:"successRate"`
			TotalPpa      float64 `json:"totalPPA"`
			Ppa           float64 `json:"ppa"`
		} `json:"rushingPlays"`
		PassingDowns struct {
			Explosiveness float64 `json:"explosiveness"`
			SuccessRate   float64 `json:"successRate"`
			Ppa           float64 `json:"ppa"`
		} `json:"passingDowns"`
		StandardDowns struct {
			Explosiveness float64 `json:"explosiveness"`
			SuccessRate   float64 `json:"successRate"`
			Ppa           float64 `json:"ppa"`
		} `json:"standardDowns"`
		OpenFieldYardsTotal   float64 `json:"openFieldYardsTotal"`
		OpenFieldYards        float64 `json:"openFieldYards"`
		SecondLevelYardsTotal float64 `json:"secondLevelYardsTotal"`
		SecondLevelYards      float64 `json:"secondLevelYards"`
		LineYardsTotal        float64 `json:"lineYardsTotal"`
		LineYards             float64 `json:"lineYards"`
		StuffRate             float64 `json:"stuffRate"`
		PowerSuccess          float64 `json:"powerSuccess"`
		Explosiveness         float64 `json:"explosiveness"`
		SuccessRate           float64 `json:"successRate"`
		TotalPpa              float64 `json:"totalPPA"`
		Ppa                   float64 `json:"ppa"`
		Drives                int     `json:"drives"`
		Plays                 int     `json:"plays"`
	} `json:"offense"`
	Defense struct {
		PassingPlays struct {
			Explosiveness float64 `json:"explosiveness"`
			SuccessRate   float64 `json:"successRate"`
			TotalPpa      float64 `json:"totalPPA"`
			Ppa           float64 `json:"ppa"`
		} `json:"passingPlays"`
		RushingPlays struct {
			Explosiveness float64 `json:"explosiveness"`
			SuccessRate   float64 `json:"successRate"`
			TotalPpa      float64 `json:"totalPPA"`
			Ppa           float64 `json:"ppa"`
		} `json:"rushingPlays"`
		PassingDowns struct {
			Explosiveness float64 `json:"explosiveness"`
			SuccessRate   float64 `json:"successRate"`
			Ppa           float64 `json:"ppa"`
		} `json:"passingDowns"`
		StandardDowns struct {
			Explosiveness float64 `json:"explosiveness"`
			SuccessRate   float64 `json:"successRate"`
			Ppa           float64 `json:"ppa"`
		} `json:"standardDowns"`
		OpenFieldYardsTotal   float64 `json:"openFieldYardsTotal"`
		OpenFieldYards        float64 `json:"openFieldYards"`
		SecondLevelYardsTotal float64 `json:"secondLevelYardsTotal"`
		SecondLevelYards      float64 `json:"secondLevelYards"`
		LineYardsTotal        float64 `json:"lineYardsTotal"`
		LineYards             float64 `json:"lineYards"`
		StuffRate             float64 `json:"stuffRate"`
		PowerSuccess          float64 `json:"powerSuccess"`
		Explosiveness         float64 `json:"explosiveness"`
		SuccessRate           float64 `json:"successRate"`
		TotalPpa              float64 `json:"totalPPA"`
		Ppa                   float64 `json:"ppa"`
		Drives                int     `json:"drives"`
		Plays                 int     `json:"plays"`
	} `json:"defense"`
}

func (f *StatsGameAdvFlat) UnmarshalJSON(data []byte) error {
	type alias StatsGameAdvRaw // alias to original struct for unmarshaling
	var raw alias

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	f.GameId = raw.GameId
	f.Season = raw.Season
	f.Week = raw.Week
	f.Team = raw.Team
	f.Opponent = raw.Opponent

	// Offense PassingPlays
	f.OffensePassingPlaysExplosiveness = raw.Offense.PassingPlays.Explosiveness
	f.OffensePassingPlaysSuccessRate = raw.Offense.PassingPlays.SuccessRate
	f.OffensePassingPlaysTotalPpa = raw.Offense.PassingPlays.TotalPpa
	f.OffensePassingPlaysPpa = raw.Offense.PassingPlays.Ppa

	// Offense RushingPlays
	f.OffenseRushingPlaysExplosiveness = raw.Offense.RushingPlays.Explosiveness
	f.OffenseRushingPlaysSuccessRate = raw.Offense.RushingPlays.SuccessRate
	f.OffenseRushingPlaysTotalPpa = raw.Offense.RushingPlays.TotalPpa
	f.OffenseRushingPlaysPpa = raw.Offense.RushingPlays.Ppa

	// Offense PassingDowns
	f.OffensePassingDownsExplosiveness = raw.Offense.PassingDowns.Explosiveness
	f.OffensePassingDownsSuccessRate = raw.Offense.PassingDowns.SuccessRate
	f.OffensePassingDownsPpa = raw.Offense.PassingDowns.Ppa

	// Offense StandardDowns
	f.OffenseStandardDownsExplosiveness = raw.Offense.StandardDowns.Explosiveness
	f.OffenseStandardDownsSuccessRate = raw.Offense.StandardDowns.SuccessRate
	f.OffenseStandardDownsPpa = raw.Offense.StandardDowns.Ppa

	// Offense other fields
	f.OffenseOpenFieldYardsTotal = raw.Offense.OpenFieldYardsTotal
	f.OffenseOpenFieldYards = raw.Offense.OpenFieldYards
	f.OffenseSecondLevelYardsTotal = raw.Offense.SecondLevelYardsTotal
	f.OffenseSecondLevelYards = raw.Offense.SecondLevelYards
	f.OffenseLineYardsTotal = raw.Offense.LineYardsTotal
	f.OffenseLineYards = raw.Offense.LineYards
	f.OffenseStuffRate = raw.Offense.StuffRate
	f.OffensePowerSuccess = raw.Offense.PowerSuccess
	f.OffenseExplosiveness = raw.Offense.Explosiveness
	f.OffenseSuccessRate = raw.Offense.SuccessRate
	f.OffenseTotalPpa = raw.Offense.TotalPpa
	f.OffensePpa = raw.Offense.Ppa
	f.OffenseDrives = raw.Offense.Drives
	f.OffensePlays = raw.Offense.Plays

	// Defense PassingPlays
	f.DefensePassingPlaysExplosiveness = raw.Defense.PassingPlays.Explosiveness
	f.DefensePassingPlaysSuccessRate = raw.Defense.PassingPlays.SuccessRate
	f.DefensePassingPlaysTotalPpa = raw.Defense.PassingPlays.TotalPpa
	f.DefensePassingPlaysPpa = raw.Defense.PassingPlays.Ppa

	// Defense RushingPlays
	f.DefenseRushingPlaysExplosiveness = raw.Defense.RushingPlays.Explosiveness
	f.DefenseRushingPlaysSuccessRate = raw.Defense.RushingPlays.SuccessRate
	f.DefenseRushingPlaysTotalPpa = raw.Defense.RushingPlays.TotalPpa
	f.DefenseRushingPlaysPpa = raw.Defense.RushingPlays.Ppa

	// Defense PassingDowns
	f.DefensePassingDownsExplosiveness = raw.Defense.PassingDowns.Explosiveness
	f.DefensePassingDownsSuccessRate = raw.Defense.PassingDowns.SuccessRate
	f.DefensePassingDownsPpa = raw.Defense.PassingDowns.Ppa

	// Defense StandardDowns
	f.DefenseStandardDownsExplosiveness = raw.Defense.StandardDowns.Explosiveness
	f.DefenseStandardDownsSuccessRate = raw.Defense.StandardDowns.SuccessRate
	f.DefenseStandardDownsPpa = raw.Defense.StandardDowns.Ppa

	// Defense other fields
	f.DefenseOpenFieldYardsTotal = raw.Defense.OpenFieldYardsTotal
	f.DefenseOpenFieldYards = raw.Defense.OpenFieldYards
	f.DefenseSecondLevelYardsTotal = raw.Defense.SecondLevelYardsTotal
	f.DefenseSecondLevelYards = raw.Defense.SecondLevelYards
	f.DefenseLineYardsTotal = raw.Defense.LineYardsTotal
	f.DefenseLineYards = raw.Defense.LineYards
	f.DefenseStuffRate = raw.Defense.StuffRate
	f.DefensePowerSuccess = raw.Defense.PowerSuccess
	f.DefenseExplosiveness = raw.Defense.Explosiveness
	f.DefenseSuccessRate = raw.Defense.SuccessRate
	f.DefenseTotalPpa = raw.Defense.TotalPpa
	f.DefensePpa = raw.Defense.Ppa
	f.DefenseDrives = raw.Defense.Drives
	f.DefensePlays = raw.Defense.Plays

	return nil
}

func FetchAndInsertGameStatsAdv() error {
	var sga StatsGameAdv
	query := fmt.Sprintf("stats/game/advanced?year=%v&week=%v&seasonType=%v", strconv.Itoa(util.SEASON), strconv.Itoa(util.WEEK), util.SEASON_TYPE)
	query = util.Trim_endpoint(query)
	conn.APICall(query, &sga)
	util.LogDBError("FetchAndInsertGameStatsAdv", util.DB.CreateInBatches(sga, 250).Error)

	return nil
}

func (StatsGameAdvFlat) TableName() string {
	return "advanced_game_stats"
}
