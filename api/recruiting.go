package api

type RecruitingTeams []RecruitingTeam
type RecruitingTeam struct {
	Year   int    `json:"year"`
	Rank   int    `json:"rank"`
	Team   string `json:"team"`
	Points int    `json:"points"`
}

type Recruits []Recruit
type Recruit struct {
	ID            string `json:"id"`
	AthleteID     string `json:"athleteId"`
	RecruitType   string `json:"recruitType"`
	Year          int    `json:"year"`
	Ranking       int    `json:"ranking"`
	Name          string `json:"name"`
	School        string `json:"school"`
	CommittedTo   string `json:"committedTo"`
	Position      string `json:"position"`
	Height        int    `json:"height"`
	Weight        int    `json:"weight"`
	Stars         int    `json:"stars"`
	Rating        int    `json:"rating"`
	City          string `json:"city"`
	StateProvince string `json:"stateProvince"`
	Country       string `json:"country"`
}

// recruiting/players?year=
// recruiting/teams?year=
