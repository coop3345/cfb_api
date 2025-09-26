package models

import (
	"cfbapi/conn"
	"cfbapi/util"
	"fmt"
	"strconv"
)

type RecruitingTeams []RecruitingTeam
type RecruitingTeam struct {
	Year   int    `json:"year"`
	Rank   int    `json:"rank"`
	Team   string `json:"team"`
	Points int    `json:"points"`
}

type Recruits []Recruit
type Recruit struct {
	ID            string `json:"id" gorm:"primaryKey"`
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

// func InsertRecruits(db *sql.DB, recruits []Recruit) error {
// 	query := `
// 	INSERT INTO recruits (
// 		id, athleteId, recruitType, year, ranking, name,
// 		school, committedTo, position, height, weight,
// 		stars, rating, city, stateProvince, country
// 	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

// 	stmt, err := db.Prepare(query)
// 	if err != nil {
// 		return fmt.Errorf("prepare failed: %w", err)
// 	}
// 	defer stmt.Close()

// 	for _, r := range recruits {
// 		_, err := stmt.Exec(
// 			r.ID, r.AthleteID, r.RecruitType, r.Year, r.Ranking, r.Name,
// 			r.School, r.CommittedTo, r.Position, r.Height, r.Weight,
// 			r.Stars, r.Rating, r.City, r.StateProvince, r.Country,
// 		)
// 		if err != nil {
// 			return fmt.Errorf("insert failed for recruit %s: %w", r.ID, err)
// 		}
// 	}

// 	return nil
// }

func FetchAndInsertRecruits(year int) error {
	var r Recruits
	query := fmt.Sprintf("recruiting/players?year=%v", strconv.Itoa(year))
	conn.APICall(query, &r)
	if err := util.DB.CreateInBatches(r, 100).Error; err != nil {
		return err
	}

	return nil
}

func FetchAndInsertRecruitingTeams(year int) error {
	var t RecruitingTeams
	query := fmt.Sprintf("recruiting/teams?year=%v", strconv.Itoa(year))
	conn.APICall(query, &t)
	if err := util.DB.CreateInBatches(t, 100).Error; err != nil {
		return err
	}

	return nil
}
