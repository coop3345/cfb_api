package weekly

import (
	"cfbapi/conn"
	"cfbapi/util"
	"encoding/json"
	"fmt"
	"strconv"
)

type SPRatings []SPRating
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

type SPRating struct {
	Year            int
	Week            int
	Team            string
	Conference      string
	Rating          float64
	Ranking         int
	SecondOrderWins float64
	Sos             float64

	// Offense fields
	OffensePace          float64
	OffenseRunRate       float64
	OffensePassingDowns  float64
	OffenseStandardDowns float64
	OffensePassing       float64
	OffenseRushing       float64
	OffenseExplosiveness float64
	OffenseSuccess       float64
	OffenseRating        float64
	OffenseRanking       int

	// Defense fields
	DefenseHavocDB         float64
	DefenseHavocFrontSeven float64
	DefenseHavocTotal      float64
	DefensePassingDowns    float64
	DefenseStandardDowns   float64
	DefensePassing         float64
	DefenseRushing         float64
	DefenseExplosiveness   float64
	DefenseSuccess         float64
	DefenseRating          float64
	DefenseRanking         int

	// Special Teams
	SpecialTeamsRating float64
}

type SRSRatings []SRS
type SRS struct {
	Year       int `json:"year"`
	Week       int
	Team       string  `json:"team"`
	Conference string  `json:"conference"`
	Division   string  `json:"division"`
	Rating     float64 `json:"rating"`
	Ranking    int     `json:"ranking"`
}

func (SRS) TableName() string {
	return "srs_ratings"
}

type FPIRatings []FPIRating
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

type FPIRating struct {
	Year                        int
	Week                        int
	Team                        string
	Conference                  string
	Fpi                         float64
	ResumeGameControl           float64
	ResumeRemainingSOS          float64
	ResumeStrengthOfSchedule    float64
	ResumeAverageWinProbability float64
	ResumeFpi                   float64
	ResumeStrengthOfRecord      float64
	EfficienciesSpecialTeams    float64
	EfficienciesDefense         float64
	EfficienciesOffense         float64
	EfficienciesOverall         float64
}

func (r *SPRatings) UnmarshalJSON(data []byte) error {
	var raw []SP
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	flat := make([]SPRating, 0, len(raw))
	for _, sp := range raw {
		flat = append(flat, SPRating{
			Year:                   sp.Year,
			Team:                   sp.Team,
			Conference:             sp.Conference,
			Rating:                 sp.Rating,
			Ranking:                sp.Ranking,
			SecondOrderWins:        sp.SecondOrderWins,
			Sos:                    sp.Sos,
			OffensePace:            sp.Offense.Pace,
			OffenseRunRate:         sp.Offense.RunRate,
			OffensePassingDowns:    sp.Offense.PassingDowns,
			OffenseStandardDowns:   sp.Offense.StandardDowns,
			OffensePassing:         sp.Offense.Passing,
			OffenseRushing:         sp.Offense.Rushing,
			OffenseExplosiveness:   sp.Offense.Explosiveness,
			OffenseSuccess:         sp.Offense.Success,
			OffenseRating:          sp.Offense.Rating,
			OffenseRanking:         sp.Offense.Ranking,
			DefenseHavocDB:         sp.Defense.Havoc.Db,
			DefenseHavocFrontSeven: sp.Defense.Havoc.FrontSeven,
			DefenseHavocTotal:      sp.Defense.Havoc.Total,
			DefensePassingDowns:    sp.Defense.PassingDowns,
			DefenseStandardDowns:   sp.Defense.StandardDowns,
			DefensePassing:         sp.Defense.Passing,
			DefenseRushing:         sp.Defense.Rushing,
			DefenseExplosiveness:   sp.Defense.Explosiveness,
			DefenseSuccess:         sp.Defense.Success,
			DefenseRating:          sp.Defense.Rating,
			DefenseRanking:         sp.Defense.Ranking,
			SpecialTeamsRating:     sp.SpecialTeams.Rating,
		})
	}
	*r = flat
	return nil
}

func (r *FPIRatings) UnmarshalJSON(data []byte) error {
	var raw []FPI
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	flat := make([]FPIRating, 0, len(raw))
	for _, fpi := range raw {
		flat = append(flat, FPIRating{
			Year:       fpi.Year,
			Team:       fpi.Team,
			Conference: fpi.Conference,
			Fpi:        fpi.Fpi,

			ResumeGameControl:           fpi.ResumeRanks.GameControl,
			ResumeRemainingSOS:          fpi.ResumeRanks.RemainingStrengthOfSchedule,
			ResumeStrengthOfSchedule:    fpi.ResumeRanks.StrengthOfSchedule,
			ResumeAverageWinProbability: fpi.ResumeRanks.AverageWinProbability,
			ResumeFpi:                   fpi.ResumeRanks.Fpi,
			ResumeStrengthOfRecord:      fpi.ResumeRanks.StrengthOfRecord,

			EfficienciesSpecialTeams: fpi.Efficiencies.SpecialTeams,
			EfficienciesDefense:      fpi.Efficiencies.Defense,
			EfficienciesOffense:      fpi.Efficiencies.Offense,
			EfficienciesOverall:      fpi.Efficiencies.Overall,
		})
	}
	*r = flat
	return nil
}

func FetchAndInsertSP() error {
	var r SPRatings
	query := fmt.Sprintf("ratings/sp?year=%v", strconv.Itoa(util.SEASON))
	conn.APICall(query, &r)

	for i := range r {
		r[i].Week = util.WEEK
	}

	util.LogDBError("FetchAndInsertSP", conn.BatchInsert(util.DB, r, 100))

	return nil
}

func FetchAndInsertSRS() error {
	var r SRSRatings
	query := fmt.Sprintf("ratings/srs?year=%v", strconv.Itoa(util.SEASON))
	conn.APICall(query, &r)

	for i := range r {
		r[i].Week = util.WEEK
	}

	util.LogDBError("FetchAndInsertSRS", conn.BatchInsert(util.DB, r, 100))

	return nil
}

func FetchAndInsertFPI() error {
	var r FPIRatings
	query := fmt.Sprintf("ratings/fpi?year=%v", strconv.Itoa(util.SEASON))
	conn.APICall(query, &r)

	for i := range r {
		r[i].Week = util.WEEK
	}

	util.LogDBError("FetchAndInsertFPI", conn.BatchInsert(util.DB, r, 100))

	return nil
}

func FetchAndInsertRatings() []error {
	var errs []error
	errs = append(errs, FetchAndInsertSP())
	errs = append(errs, FetchAndInsertSRS())
	errs = append(errs, FetchAndInsertFPI())

	return errs
}
