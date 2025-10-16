package models

import (
	"cfbapi/conn"
	"cfbapi/util"
	"fmt"
	"log"
	"strconv"
)

type RecruitingTeams []RecruitingTeam
type RecruitingTeam struct {
	Year   int     `json:"year"`
	Rank   int     `json:"rank"`
	Team   string  `json:"team"`
	Points float64 `json:"points"`
}

type Recruits []Recruit
type Recruit struct {
	RecruitID     string  `json:"id"`
	AthleteID     string  `json:"athleteId"`
	RecruitType   string  `json:"recruitType"`
	Year          int     `json:"year"`
	Ranking       int     `json:"ranking"`
	Name          string  `json:"name"`
	School        string  `json:"school"`
	CommittedTo   string  `json:"committedTo"`
	Position      string  `json:"position"`
	Height        int     `json:"height"`
	Weight        int     `json:"weight"`
	Stars         int     `json:"stars"`
	Rating        float64 `json:"rating"`
	City          string  `json:"city"`
	StateProvince string  `json:"stateProvince"`
	Country       string  `json:"country"`
}

func FetchAndInsertRecruits(year int) error {
	var r Recruits
	query := fmt.Sprintf("recruiting/players?year=%v", strconv.Itoa(year))
	conn.APICall(query, &r)
	if err := conn.BatchInsert(util.CONFIG.CONNECTIONS.DB, r, 100); err != nil {
		return err
	}
	log.Println("Inserted", len(r), "recruit records.")

	return nil
}

func FetchAndInsertRecruitingTeams(year int) error {
	var t RecruitingTeams
	query := fmt.Sprintf("recruiting/teams?year=%v", strconv.Itoa(year))
	conn.APICall(query, &t)
	if err := conn.BatchInsert(util.CONFIG.CONNECTIONS.DB, t, 100); err != nil {
		return err
	}
	log.Println("Inserted", len(t), "team recruiting records.")

	return nil
}
