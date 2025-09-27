package weekly

import (
	"cfbapi/conn"
	"cfbapi/util"
	"fmt"
	"strconv"
)

type GameWeather struct {
	GameID               int     `json:"id" gorm:"primaryKey"`
	Season               int     `json:"season"`
	Week                 int     `json:"week"`
	SeasonType           string  `json:"seasonType"`
	StartTime            string  `json:"startTime"`
	GameIndoors          bool    `json:"gameIndoors"`
	HomeTeam             string  `json:"homeTeam"`
	HomeConference       string  `json:"homeConference"`
	AwayTeam             string  `json:"awayTeam"`
	AwayConference       string  `json:"awayConference"`
	VenueId              int     `json:"venueId"`
	Venue                string  `json:"venue"`
	Temperature          float64 `json:"temperature"`
	DewPoint             float64 `json:"dewPoint"`
	Humidity             float64 `json:"humidity"`
	Precipitation        float64 `json:"precipitation"`
	Snowfall             float64 `json:"snowfall"`
	WindDirection        float64 `json:"windDirection"`
	WindSpeed            float64 `json:"windSpeed"`
	Pressure             float64 `json:"pressure"`
	WeatherConditionCode int     `json:"weatherConditionCode"`
	WeatherCondition     string  `json:"weatherCondition"`
}

func (GameWeather) TableName() string {
	return "game_weather"
}

func FetchAndInsertGameWeather() error {
	var gameWeather []GameWeather
	query := fmt.Sprintf("games/weather?year=%v&week=%v&seasonType=%v", strconv.Itoa(util.SEASON), strconv.Itoa(util.WEEK), util.SEASON_TYPE)
	query = util.Trim_endpoint(query)
	conn.APICall(query, &gameWeather)
	util.LogDBError("FetchAndInsertGameWeather", util.DB.CreateInBatches(gameWeather, 1).Error)

	return nil
}
