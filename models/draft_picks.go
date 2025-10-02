package models

import (
	"cfbapi/conn"
	"cfbapi/util"
	"fmt"
	"strconv"
)

type DraftPicks []DraftPick
type DraftPick struct {
	CollegeAthleteId        int     `json:"collegeAthleteId"`
	NflAthleteId            int     `json:"nflAthleteId"`
	CollegeId               int     `json:"collegeId"`
	CollegeTeam             string  `json:"collegeTeam"`
	CollegeConference       string  `json:"collegeConference"`
	NflTeamId               int     `json:"nflTeamId"`
	NflTeam                 string  `json:"nflTeam"`
	Year                    int     `json:"year"`
	Overall                 int     `json:"overall"`
	Round                   int     `json:"round"`
	Pick                    int     `json:"pick"`
	Name                    string  `json:"name"`
	Position                string  `json:"position"`
	Height                  int     `json:"height"`
	Weight                  int     `json:"weight"`
	PreDraftRanking         float64 `json:"preDraftRanking"`
	PreDraftPositionRanking float64 `json:"preDraftPositionRanking"`
	PreDraftGrade           float64 `json:"preDraftGrade"`
}

func FetchAndInsertDraftPicks(year int) error {
	var dp DraftPicks
	query := fmt.Sprintf("draft/picks?year=%v", strconv.Itoa(year))
	conn.APICall(query, &dp)
	if err := conn.BatchInsert(util.DB, dp, 100); err != nil {
		return err
	}

	return nil
}
