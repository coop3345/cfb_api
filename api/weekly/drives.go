package weekly

type Drives []Drive
type Drive struct {
	Offense           string `json:"offense"`
	OffenseConference string `json:"offenseConference"`
	Defense           string `json:"defense"`
	DefenseConference string `json:"defenseConference"`
	GameId            int    `json:"gameId"`
	Id                string `json:"id"`
	DriveNumber       int    `json:"driveNumber"`
	Scoring           bool   `json:"scoring"`
	StartPeriod       int    `json:"startPeriod"`
	StartYardline     int    `json:"startYardline"`
	StartYardsToGoal  int    `json:"startYardsToGoal"`
	StartTime         struct {
		Seconds int `json:"seconds"`
		Minutes int `json:"minutes"`
	} `json:"startTime"`
	EndPeriod      int `json:"endPeriod"`
	EndYardline    int `json:"endYardline"`
	EndYardsToGoal int `json:"endYardsToGoal"`
	EndTime        struct {
		Seconds int `json:"seconds"`
		Minutes int `json:"minutes"`
	} `json:"endTime"`
	Elapsed struct {
		Seconds int `json:"seconds"`
		Minutes int `json:"minutes"`
	} `json:"elapsed"`
	Plays             int    `json:"plays"`
	Yards             int    `json:"yards"`
	DriveResult       string `json:"driveResult"`
	IsHomeOffense     bool   `json:"isHomeOffense"`
	StartOffenseScore int    `json:"startOffenseScore"`
	StartDefenseScore int    `json:"startDefenseScore"`
	EndOffenseScore   int    `json:"endOffenseScore"`
	EndDefenseScore   int    `json:"endDefenseScore"`
}
