package seasonal

import (
	"cfbapi/conn"
	"cfbapi/util"
	"encoding/json"
	"fmt"
	"strconv"
)

type Coaches []Coach
type Coach struct {
	FirstName      string  `json:"firstName" gorm:"primaryKey;size:40"`
	LastName       string  `json:"lastName" gorm:"primaryKey;size:40"`
	HireDate       string  `json:"hireDate"`
	School         string  `json:"school" gorm:"primaryKey;size:100"`
	Year           int     `json:"year" gorm:"primaryKey"`
	Games          int     `json:"games"`
	Wins           int     `json:"wins"`
	Losses         int     `json:"losses"`
	Ties           int     `json:"ties"`
	PreseasonRank  int     `json:"preseasonRank"`
	PostseasonRank int     `json:"postseasonRank"`
	Srs            float64 `json:"srs"`
	SpOverall      float64 `json:"spOverall"`
	SpOffense      float64 `json:"spOffense"`
	SpDefense      float64 `json:"spDefense"`
}

func (c *Coaches) UnmarshalJSON(data []byte) error {
	var raw []struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		HireDate  string `json:"hireDate"`
		Seasons   []struct {
			School         string  `json:"school"`
			Year           int     `json:"year"`
			Games          int     `json:"games"`
			Wins           int     `json:"wins"`
			Losses         int     `json:"losses"`
			Ties           int     `json:"ties"`
			PreseasonRank  int     `json:"preseasonRank"`
			PostseasonRank int     `json:"postseasonRank"`
			Srs            float64 `json:"srs"`
			SpOverall      float64 `json:"spOverall"`
			SpOffense      float64 `json:"spOffense"`
			SpDefense      float64 `json:"spDefense"`
		} `json:"seasons"`
	}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	var flattened Coaches
	for _, coach := range raw {
		for _, season := range coach.Seasons {
			flattened = append(flattened, Coach{
				FirstName:      coach.FirstName,
				LastName:       coach.LastName,
				HireDate:       coach.HireDate,
				School:         season.School,
				Year:           season.Year,
				Games:          season.Games,
				Wins:           season.Wins,
				Losses:         season.Losses,
				Ties:           season.Ties,
				PreseasonRank:  season.PreseasonRank,
				PostseasonRank: season.PostseasonRank,
				Srs:            season.Srs,
				SpOverall:      season.SpOverall,
				SpOffense:      season.SpOffense,
				SpDefense:      season.SpDefense,
			})
		}
	}

	*c = flattened
	return nil
}

func FetchAndInsertCoaches() error {
	var coaches Coaches
	query := fmt.Sprintf("coaches?year=%v", strconv.Itoa(util.SEASON))
	conn.APICall(query, &coaches)

	util.LogDBError("FetchAndInsertCoaches", conn.BatchInsert(util.CONFIG.CONNECTIONS.DB, coaches, 1))

	return nil
}
