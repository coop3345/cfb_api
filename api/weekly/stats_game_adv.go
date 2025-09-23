package weekly

type StatsGameAdv struct {
	GameId   int    `json:"gameId"`
	Season   int    `json:"season"`
	Week     int    `json:"week"`
	Team     string `json:"team"`
	Opponent string `json:"opponent"`
	Offense  struct {
		PassingPlays struct {
			Explosiveness float64 `json:"explosiveness"`
			SuccessRate   float64 `json:"successRate"`
			TotalPpa      float64 `json:"totalPPA"`
			Ppa           float64 `json:"ppa"`
		} `json:"passingPlays"`
		RushingPlays struct {
			Explosiveness float64 `json:"explosiveness"`
			SuccessRate   float64 `json:"successRate"`
			TotalPpa      float64 `json:"totalPPA"`
			Ppa           float64 `json:"ppa"`
		} `json:"rushingPlays"`
		PassingDowns struct {
			Explosiveness float64 `json:"explosiveness"`
			SuccessRate   float64 `json:"successRate"`
			Ppa           float64 `json:"ppa"`
		} `json:"passingDowns"`
		StandardDowns struct {
			Explosiveness float64 `json:"explosiveness"`
			SuccessRate   float64 `json:"successRate"`
			Ppa           float64 `json:"ppa"`
		} `json:"standardDowns"`
		OpenFieldYardsTotal   float64 `json:"openFieldYardsTotal"`
		OpenFieldYards        float64 `json:"openFieldYards"`
		SecondLevelYardsTotal float64 `json:"secondLevelYardsTotal"`
		SecondLevelYards      float64 `json:"secondLevelYards"`
		LineYardsTotal        float64 `json:"lineYardsTotal"`
		LineYards             float64 `json:"lineYards"`
		StuffRate             float64 `json:"stuffRate"`
		PowerSuccess          float64 `json:"powerSuccess"`
		Explosiveness         float64 `json:"explosiveness"`
		SuccessRate           float64 `json:"successRate"`
		TotalPpa              float64 `json:"totalPPA"`
		Ppa                   float64 `json:"ppa"`
		Drives                int     `json:"drives"`
		Plays                 int     `json:"plays"`
	} `json:"offense"`
	Defense struct {
		PassingPlays struct {
			Explosiveness float64 `json:"explosiveness"`
			SuccessRate   float64 `json:"successRate"`
			TotalPpa      float64 `json:"totalPPA"`
			Ppa           float64 `json:"ppa"`
		} `json:"passingPlays"`
		RushingPlays struct {
			Explosiveness float64 `json:"explosiveness"`
			SuccessRate   float64 `json:"successRate"`
			TotalPpa      float64 `json:"totalPPA"`
			Ppa           float64 `json:"ppa"`
		} `json:"rushingPlays"`
		PassingDowns struct {
			Explosiveness float64 `json:"explosiveness"`
			SuccessRate   float64 `json:"successRate"`
			Ppa           float64 `json:"ppa"`
		} `json:"passingDowns"`
		StandardDowns struct {
			Explosiveness float64 `json:"explosiveness"`
			SuccessRate   float64 `json:"successRate"`
			Ppa           float64 `json:"ppa"`
		} `json:"standardDowns"`
		OpenFieldYardsTotal   float64 `json:"openFieldYardsTotal"`
		OpenFieldYards        float64 `json:"openFieldYards"`
		SecondLevelYardsTotal float64 `json:"secondLevelYardsTotal"`
		SecondLevelYards      float64 `json:"secondLevelYards"`
		LineYardsTotal        float64 `json:"lineYardsTotal"`
		LineYards             float64 `json:"lineYards"`
		StuffRate             float64 `json:"stuffRate"`
		PowerSuccess          float64 `json:"powerSuccess"`
		Explosiveness         float64 `json:"explosiveness"`
		SuccessRate           float64 `json:"successRate"`
		TotalPpa              float64 `json:"totalPPA"`
		Ppa                   float64 `json:"ppa"`
		Drives                int     `json:"drives"`
		Plays                 int     `json:"plays"`
	} `json:"defense"`
}
