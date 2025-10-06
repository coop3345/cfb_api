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
	util.InitLogger()
	util.Setup()
	util.CONFIG.CONNECTIONS.DB, _ = conn.InitDB()
	models.Migrate_Model()

	models.FetchAndInsertConferences()

	for y := util.CONFIG.RUN_PARAMS.START_SEASON; y <= util.CONFIG.RUN_PARAMS.END_SEASON; y++ {
		util.SEASON = y
		// get entire season backfill
		if util.CONFIG.RUN_PARAMS.GET_SEASON {
			get_season(y)
		} else if util.CONFIG.RUN_PARAMS.GET_WEEKLY { // get just 1 week - iterate through calendar to find most recent week
			cal, _ := seasonal.FetchAndInsertCalendar()
			for _, week := range cal {
				weekEnd := week.EndDate.Unix()
				now := time.Now().Unix()

				if weekEnd < now {
					util.WEEK, util.SEASON_TYPE = week.Week, week.SeasonType
					continue
				} else {
					weekly.RANK_WEEK, weekly.RANK_SEASON_TYPE = week.Week, week.SeasonType
					break
				}
			}
			get_week(y, util.WEEK)
		}
		if util.CONFIG.RUN_PARAMS.GET_OFFSEASON {
			models.FetchAndInsertDraftPicks(y)
			models.FetchAndInsertRecruitingTeams(y)
			models.FetchAndInsertRecruits(y)
		}
	}
	if util.CONFIG.RUN_PARAMS.GET_ONE_OFFS {
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

	if !util.CONFIG.RUN_PARAMS.GET_WEEKLY {
		return
	}
	for _, week := range cal {
		util.WEEK, weekly.RANK_WEEK = week.Week, week.Week
		util.SEASON_TYPE, weekly.RANK_SEASON_TYPE = week.SeasonType, week.SeasonType

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

	for _, game := range weekly.GAMES {
		if util.Contains(util.CONFIG.RUN_PARAMS.PSCD, game.AwayClassification) || util.Contains(util.CONFIG.RUN_PARAMS.PSCD, game.HomeClassification) {
			weekly.FetchAndInsertPlayStatsGame(game.Id)
		}
	}
}

// func backfill_play_stats(year int) {
// 	util.SEASON = year
// 	cal, _ := seasonal.FetchAndInsertCalendar()
// 	for _, week := range cal {
// 		util.WEEK = week.Week
// 		util.SEASON_TYPE = week.SeasonType
// 		if week.EndDate.Unix() > time.Now().Unix() {
// 			fmt.Printf("Week (%v) not yet completed in Season - %v", week, year)
// 			break
// 		} else if week.Week == 5 && year == 2025 {
// 			break
// 		}

// 		query := fmt.Sprintf("games?year=%v&week=%v&seasonType=%v", strconv.Itoa(util.SEASON), strconv.Itoa(util.WEEK), util.SEASON_TYPE)
// 		conn.APICall(query, &weekly.GAMES)

// 		for _, game := range weekly.GAMES {
// 			if util.Contains(util.PSCD, game.AwayClassification) || util.Contains(util.PSCD, game.HomeClassification) {
// 				weekly.FetchAndInsertPlayStatsGame(game.Id)
// 			}
// 		}
// 	}
// }
