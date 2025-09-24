package weekly

import "encoding/json"

type Rankings []RankingFlat

// todo: flatten
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
