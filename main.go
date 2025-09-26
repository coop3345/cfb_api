package main

import (
	"cfbapi/conn"
	"cfbapi/models"
	"cfbapi/models/seasonal"
	"cfbapi/models/weekly"
	"cfbapi/util"
	"fmt"
	"time"
)

func main() {
	util.DB, _ = conn.InitDB()

	models.FetchAndInsertConferences()
	for y := util.START_SEASON; y <= util.END_SEASON; y++ {
		if util.GET_SEASON {
			get_season(y)
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

	for _, con := range models.CONFERENCES {
		weekly.FetchAndInsertPlayStats(con.Name)
	}
}
