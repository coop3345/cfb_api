package main

import (
	"cfbapi/conn"
	"cfbapi/models/seasonal"
	"cfbapi/util"
	"fmt"
	"time"
)

func main() {
	util.DB, _ = conn.InitDB()

	if util.GET_SEASON {
		for y := util.START_SEASON; y <= util.END_SEASON; y++ {
			get_season(y)
		}
	}

	if util.GET_PICKS {

	}

	if util.GET_ONE_OFFS {

	}
}

func get_season(year int) {
	cal, _ := seasonal.FetchAndInsertCalendar()
	seasonal.FetchAndInsertCoaches()

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
}
