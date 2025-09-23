package weekly

//todo flatten
type GamePlayerStat struct {
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
