package weekly

import (
	"cfbapi/conn"
	"cfbapi/util"
	"fmt"
	"log"
	"strconv"
)

var GAMES Games

type Games []Game
type Game struct {
	Id                         int     `json:"id" gorm:"primaryKey"`
	Season                     int     `json:"season"`
	Week                       int     `json:"week"`
	SeasonType                 string  `json:"seasonType"`
	StartDate                  string  `json:"startDate"`
	StartTimeTBD               bool    `json:"startTimeTBD"`
	Completed                  bool    `json:"completed"`
	NeutralSite                bool    `json:"neutralSite"`
	ConferenceGame             bool    `json:"conferenceGame"`
	Attendance                 int     `json:"attendance"`
	VenueId                    int     `json:"venueId"`
	Venue                      string  `json:"venue"`
	HomeId                     int     `json:"homeId"`
	HomeTeam                   string  `json:"homeTeam"`
	HomeConference             string  `json:"homeConference"`
	HomeClassification         string  `json:"homeClassification"`
	HomePoints                 int     `json:"homePoints"`
	HomeLineScores             string  `json:"homeLineScores" gorm.type:"NVARCHAR(MAX)"`
	HomePostgameWinProbability float64 `json:"homePostgameWinProbability"`
	HomePregameElo             float64 `json:"homePregameElo"`
	HomePostgameElo            float64 `json:"homePostgameElo"`
	AwayId                     int     `json:"awayId"`
	AwayTeam                   string  `json:"awayTeam"`
	AwayConference             string  `json:"awayConference"`
	AwayClassification         string  `json:"awayClassification"`
	AwayPoints                 int     `json:"awayPoints"`
	AwayLineScores             string  `json:"awayLineScores" gorm.type:"NVARCHAR(MAX)"`
	AwayPostgameWinProbability float64 `json:"awayPostgameWinProbability"`
	AwayPregameElo             float64 `json:"awayPregameElo"`
	AwayPostgameElo            float64 `json:"awayPostgameElo"`
	ExcitementIndex            float64 `json:"excitementIndex"`
	Highlights                 string  `json:"highlights"`
	Notes                      string  `json:"notes"`
}

func FetchAndInsertGames() error {
	query := fmt.Sprintf("games?year=%v&week=%v&seasonType=%v", strconv.Itoa(util.SEASON), strconv.Itoa(util.WEEK), util.SEASON_TYPE)
	query = util.Trim_endpoint(query)

	conn.APICall(query, &GAMES)
	util.LogDBError("FetchAndInsertGames", conn.BatchInsert(util.CONFIG.CONNECTIONS.DB, GAMES, 1))
	log.Println("Inserted", len(GAMES), "games.")

	return nil
}
