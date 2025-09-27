package models

import (
	"cfbapi/models/seasonal"
	"cfbapi/models/weekly"
	"cfbapi/util"
)

func Migrate_Model() {
	util.DB.AutoMigrate(
		&Recruit{},
		&Conference{},
		&DraftPick{},
		&PlayStatType{},
		&PlayType{},
		&RecruitingTeam{},
		&Venue{},
		&seasonal.Week{},
		&seasonal.Coach{},
		&seasonal.PlayerPortalEntry{},
		&seasonal.PlayerUsage{},
		&seasonal.Roster{},
		&seasonal.Talent{},
		&seasonal.Team{},
		&weekly.Drive{},
		&weekly.GamePlayerStat{},
		&weekly.GameTeamStat{},
		&weekly.Game{},
		&weekly.GameWeather{},
		&weekly.Play{},
		&weekly.PlayStat{},
		&weekly.RankingFlat{},
		&weekly.SPRating{},
		&weekly.SRS{},
		&weekly.FPIRating{},
		&weekly.StatsGameAdvFlat{})
}
