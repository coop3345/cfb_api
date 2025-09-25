package models

import (
	"cfbapi/conn"
	"cfbapi/util"
	"encoding/json"
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

func FetchAndInsertDraftPicks() error {
	var dp DraftPicks
	query := fmt.Sprintf("draft/picks?year=%v", strconv.Itoa(util.SEASON))

	b, _ := conn.APICall(query)
	if err := json.Unmarshal(b, &dp); err != nil {
		panic(err)
	}
	if err := util.DB.CreateInBatches(dp, 100).Error; err != nil {
		return err
	}

	return nil
}
