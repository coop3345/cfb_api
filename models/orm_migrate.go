package models

import (
	"cfbapi/models/seasonal"
	"cfbapi/models/weekly"
	"cfbapi/util"
	"log"
)

func Migrate_Model() error {
	err := util.CONFIG.CONNECTIONS.DB.AutoMigrate(
		&Recruit{},
		&Conference{},
		&Info{},
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
		&weekly.StatsGameAdvFlat{},
		&weekly.GameLines{})

	if err != nil {
		log.Fatalf("Failed to migrate models: %v", err)
	}

	return err
}
