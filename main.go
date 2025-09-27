package main

import (
	"cfbapi/conn"
	"cfbapi/models"
	"cfbapi/models/seasonal"
	"cfbapi/models/weekly"
	"cfbapi/util"
	"fmt"
	"strconv"
	"time"
)

func main() {
	util.InitLogger()
	util.DB, _ = conn.InitDB()
	models.Migrate_Model()

	models.FetchAndInsertConferences()

	for y := util.START_SEASON; y <= util.END_SEASON; y++ {
		if util.GET_SEASON {
			util.SEASON = y
			get_season(y)
		} else if util.GET_WEEKLY {
			cal, _ := seasonal.FetchAndInsertCalendar()
			for _, week := range cal {
				if week.EndDate.Unix() < time.Now().Unix() {
					util.WEEK, util.SEASON_TYPE = week.Week, week.SeasonType
					continue
				} else {
					get_week(y, util.WEEK)
					break
				}
			}
			get_week(y, util.WEEK)
		}
		if util.GET_OFFSEASON {
			models.FetchAndInsertDraftPicks(y)
			models.FetchAndInsertRecruitingTeams(y)
			models.FetchAndInsertRecruits(y)
		}
	}
	if util.GET_ONE_OFFS {
		models.FetchAndInsertPlayStatTypes()
		models.FetchAndInsertPlayTypes()
		models.FetchAndInsertVenues()
	}
}

func get_season(year int) {
	cal, _ := seasonal.FetchAndInsertCalendar()
	seasonal.FetchAndInsertCoaches()
	seasonal.FetchAndInsertPortal()
	seasonal.FetchAndInsertRosters()
	seasonal.FetchAndInsertTalent()
	seasonal.FetchAndInsertTeams()
	seasonal.FetchAndInsertPlayerUsage()

	if !util.GET_WEEKLY {
		return
	}
	for _, week := range cal {
		util.WEEK = week.Week
		util.SEASON_TYPE = week.SeasonType
		if week.EndDate.Unix() > time.Now().Unix() {
			fmt.Printf("Week (%v) not yet completed in Season - %v", week, year)
			break
		}
		get_week(util.SEASON, util.WEEK)
	}
}

func get_week(year int, week int) {
	print(year, ":", week)
	weekly.FetchAndInsertGames()
	weekly.FetchAndInsertDrives()
	weekly.FetchAndInsertGamePlayerStats()
	weekly.FetchAndInsertGameTeamStats()
	weekly.FetchAndInsertGameWeather()
	weekly.FetchAndInsertRankings()
	weekly.FetchAndInsertRatings()
	weekly.FetchAndInsertGameStatsAdv()
	weekly.FetchAndInsertPlays()

	for _, con := range models.CONFERENCES {
		if util.Contains(util.PSCD, con.Classification) {
			weekly.FetchAndInsertPlayStats(con.Name)
		}
	}

	fmt.Printf("Fetching lower div play stats vs D1 teams. Game Count = %v", strconv.Itoa(len(weekly.LOWERDIVGAMES)))
	for gameId, team := range weekly.LOWERDIVGAMES {
		weekly.FetchAndInsertPlayStatsGame(gameId, team)
	}

}
