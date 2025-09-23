package weekly

type Rankings []Ranking

// todo: flatten
type Ranking struct {
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
