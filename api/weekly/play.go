package weekly

type Plays []Play
type Play struct {
	Id                string `json:"id"`
	DriveId           string `json:"driveId"`
	GameId            int    `json:"gameId"`
	DriveNumber       int    `json:"driveNumber"`
	PlayNumber        int    `json:"playNumber"`
	Offense           string `json:"offense"`
	OffenseConference string `json:"offenseConference"`
	OffenseScore      int    `json:"offenseScore"`
	Defense           string `json:"defense"`
	Home              string `json:"home"`
	Away              string `json:"away"`
	DefenseConference string `json:"defenseConference"`
	DefenseScore      int    `json:"defenseScore"`
	Period            int    `json:"period"`
	Clock             struct {
		Seconds int `json:"seconds"`
		Minutes int `json:"minutes"`
	} `json:"clock"`
	OffenseTimeouts int     `json:"offenseTimeouts"`
	DefenseTimeouts int     `json:"defenseTimeouts"`
	Yardline        int     `json:"yardline"`
	YardsToGoal     int     `json:"yardsToGoal"`
	Down            int     `json:"down"`
	Distance        int     `json:"distance"`
	YardsGained     int     `json:"yardsGained"`
	Scoring         bool    `json:"scoring"`
	PlayType        string  `json:"playType"`
	PlayText        string  `json:"playText"`
	Ppa             float64 `json:"ppa"`
	Wallclock       string  `json:"wallclock"`
}

type PlayStat struct {
	GameId        int    `json:"gameId"`
	Season        int    `json:"season"`
	Week          int    `json:"week"`
	Team          string `json:"team"`
	Conference    string `json:"conference"`
	Opponent      string `json:"opponent"`
	TeamScore     int    `json:"teamScore"`
	OpponentScore int    `json:"opponentScore"`
	DriveId       string `json:"driveId"`
	PlayId        string `json:"playId"`
	Period        int    `json:"period"`
	Clock         struct {
		Seconds int `json:"seconds"`
		Minutes int `json:"minutes"`
	} `json:"clock"`
	YardsToGoal int    `json:"yardsToGoal"`
	Down        int    `json:"down"`
	Distance    int    `json:"distance"`
	AthleteId   string `json:"athleteId"`
	AthleteName string `json:"athleteName"`
	StatType    string `json:"statType"`
	Stat        int    `json:"stat"`
}
