package weekly

var GameEndpoint = "/games/teams"

type GameTeamStats []GameTeamStat
type GameTeamStat struct {
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
