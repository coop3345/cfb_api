package weekly

import (
	"cfbapi/conn"
	"cfbapi/util"
	"fmt"
	"strconv"

	"gorm.io/datatypes"
)

var GAMES map[string]int
var LOWERDIVGAMES map[int]string

type Games []Game
type Game struct {
	Id                         int            `json:"id" gorm:"primaryKey"`
	Season                     int            `json:"season"`
	Week                       int            `json:"week"`
	SeasonType                 string         `json:"seasonType"`
	StartDate                  string         `json:"startDate"`
	StartTimeTBD               bool           `json:"startTimeTBD"`
	Completed                  bool           `json:"completed"`
	NeutralSite                bool           `json:"neutralSite"`
	ConferenceGame             bool           `json:"conferenceGame"`
	Attendance                 int            `json:"attendance"`
	VenueId                    int            `json:"venueId"`
	Venue                      string         `json:"venue"`
	HomeId                     int            `json:"homeId"`
	HomeTeam                   string         `json:"homeTeam"`
	HomeConference             string         `json:"homeConference"`
	HomeClassification         string         `json:"homeClassification"`
	HomePoints                 int            `json:"homePoints"`
	HomeLineScores             datatypes.JSON `json:"homeLineScores"`
	HomePostgameWinProbability float64        `json:"homePostgameWinProbability"`
	HomePregameElo             float64        `json:"homePregameElo"`
	HomePostgameElo            float64        `json:"homePostgameElo"`
	AwayId                     int            `json:"awayId"`
	AwayTeam                   string         `json:"awayTeam"`
	AwayConference             string         `json:"awayConference"`
	AwayClassification         string         `json:"awayClassification"`
	AwayPoints                 int            `json:"awayPoints"`
	AwayLineScores             datatypes.JSON `json:"awayLineScores"`
	AwayPostgameWinProbability float64        `json:"awayPostgameWinProbability"`
	AwayPregameElo             float64        `json:"awayPregameElo"`
	AwayPostgameElo            float64        `json:"awayPostgameElo"`
	ExcitementIndex            float64        `json:"excitementIndex"`
	Highlights                 string         `json:"highlights"`
	Notes                      string         `json:"notes"`
}

func FetchAndInsertGames() error {
	var games Games
	query := fmt.Sprintf("games?year=%v&week=%v&seasonType=%v", strconv.Itoa(util.SEASON), strconv.Itoa(util.WEEK), util.SEASON_TYPE)
	query = util.Trim_endpoint(query)
	conn.APICall(query, &games)
	if err := util.DB.CreateInBatches(games, 100).Error; err != nil {
		return err
	}

	GAMES = make(map[string]int)
	LOWERDIVGAMES = make(map[int]string)
	for _, game := range games {
		GAMES[game.AwayTeam] = game.Id
		GAMES[game.HomeTeam] = game.Id

		if (util.Contains(util.PSCD, game.HomeClassification) && !util.Contains(util.PSCD, game.AwayClassification)) || (!util.Contains(util.PSCD, game.HomeClassification) && util.Contains(util.PSCD, game.AwayClassification)) {
			if !util.Contains(util.PSCD, game.AwayClassification) {
				LOWERDIVGAMES[game.Id] = game.AwayTeam
			} else if !util.Contains(util.PSCD, game.HomeClassification) {
				LOWERDIVGAMES[game.Id] = game.HomeTeam
			}
		}
	}

	return nil
}
