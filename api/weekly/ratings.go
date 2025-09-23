package weekly

type SPRatings []SP
type SP struct {
	Year            int     `json:"year"`
	Team            string  `json:"team"`
	Conference      string  `json:"conference"`
	Rating          float64 `json:"rating"`
	Ranking         int     `json:"ranking"`
	SecondOrderWins float64 `json:"secondOrderWins"`
	Sos             float64 `json:"sos"`
	Offense         struct {
		Pace          float64 `json:"pace"`
		RunRate       float64 `json:"runRate"`
		PassingDowns  float64 `json:"passingDowns"`
		StandardDowns float64 `json:"standardDowns"`
		Passing       float64 `json:"passing"`
		Rushing       float64 `json:"rushing"`
		Explosiveness float64 `json:"explosiveness"`
		Success       float64 `json:"success"`
		Rating        float64 `json:"rating"`
		Ranking       int     `json:"ranking"`
	} `json:"offense"`
	Defense struct {
		Havoc struct {
			Db         float64 `json:"db"`
			FrontSeven float64 `json:"frontSeven"`
			Total      float64 `json:"total"`
		} `json:"havoc"`
		PassingDowns  float64 `json:"passingDowns"`
		StandardDowns float64 `json:"standardDowns"`
		Passing       float64 `json:"passing"`
		Rushing       float64 `json:"rushing"`
		Explosiveness float64 `json:"explosiveness"`
		Success       float64 `json:"success"`
		Rating        float64 `json:"rating"`
		Ranking       int     `json:"ranking"`
	} `json:"defense"`
	SpecialTeams struct {
		Rating float64 `json:"rating"`
	} `json:"specialTeams"`
}

type SRSRatings []SRS
type SRS struct {
	Year       int     `json:"year"`
	Team       string  `json:"team"`
	Conference string  `json:"conference"`
	Division   string  `json:"division"`
	Rating     float64 `json:"rating"`
	Ranking    int     `json:"ranking"`
}

type FPIRatings []FPI
type FPI struct {
	Year        int     `json:"year"`
	Team        string  `json:"team"`
	Conference  string  `json:"conference"`
	Fpi         float64 `json:"fpi"`
	ResumeRanks struct {
		GameControl                 float64 `json:"gameControl"`
		RemainingStrengthOfSchedule float64 `json:"remainingStrengthOfSchedule"`
		StrengthOfSchedule          float64 `json:"strengthOfSchedule"`
		AverageWinProbability       float64 `json:"averageWinProbability"`
		Fpi                         float64 `json:"fpi"`
		StrengthOfRecord            float64 `json:"strengthOfRecord"`
	} `json:"resumeRanks"`
	Efficiencies struct {
		SpecialTeams float64 `json:"specialTeams"`
		Defense      float64 `json:"defense"`
		Offense      float64 `json:"offense"`
		Overall      float64 `json:"overall"`
	} `json:"efficiencies"`
}
