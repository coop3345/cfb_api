package weekly

import (
	"cfbapi/conn"
	"cfbapi/util"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

var RANK_SEASON_TYPE string
var RANK_WEEK int

type Rankings []RankingFlat
type RankingRaw struct {
	Season     int    `json:"season"`
	SeasonType string `json:"seasonType"`
	Week       int    `json:"week"`
	Polls      []struct {
		Poll  string `json:"poll"`
		Ranks []struct {
			Rank            int    `json:"rank"`
			TeamId          int    `json:"teamId"`
			School          string `json:"school"`
			Conference      string `json:"conference"`
			FirstPlaceVotes int    `json:"firstPlaceVotes"`
			Points          int    `json:"points"`
		} `json:"ranks"`
	} `json:"polls"`
}

type RankingFlat struct {
	Season          int
	SeasonType      string
	Week            int
	Poll            string
	Rank            int
	TeamId          int
	School          string
	Conference      string
	FirstPlaceVotes int
	Points          int
}

func (RankingFlat) TableName() string {
	return "rankings"
}

func (r *Rankings) UnmarshalJSON(data []byte) error {
	// First unmarshal into the nested structure slice
	var nested []RankingRaw

	if err := json.Unmarshal(data, &nested); err != nil {
		return err
	}

	var flatEntries []RankingFlat
	for _, ranking := range nested {
		for _, poll := range ranking.Polls {
			for _, rank := range poll.Ranks {
				flatEntries = append(flatEntries, RankingFlat{
					Season:          ranking.Season,
					SeasonType:      ranking.SeasonType,
					Week:            ranking.Week,
					Poll:            poll.Poll,
					Rank:            rank.Rank,
					TeamId:          rank.TeamId,
					School:          rank.School,
					Conference:      rank.Conference,
					FirstPlaceVotes: rank.FirstPlaceVotes,
					Points:          rank.Points,
				})
			}
		}
	}

	*r = flatEntries
	return nil
}

func FetchAndInsertRankings() error {
	var r Rankings
	query := fmt.Sprintf("rankings?year=%v&week=%v&seasonType=%v", strconv.Itoa(util.SEASON), RANK_WEEK, RANK_SEASON_TYPE)
	query = util.Trim_endpoint(query)

	conn.APICall(query, &r)
	util.LogDBError("FetchAndInsertRankings", conn.BatchInsert(util.CONFIG.CONNECTIONS.DB, r, 100))
	log.Println("Inserted", len(r), "rankings.")

	return nil
}
