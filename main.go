package main

import (
	"cfbapi/api/seasonal"
	"cfbapi/util"
	"time"
)

func main() {
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
	cal := seasonal.GetCalendar()
	seasonal.GetCoaches()

	if !util.GET_WEEKLY {
		return
	}
	for _, week := range cal {
		util.WEEK = week.Week
		if week.EndDate.Unix() > time.Now().Unix() {
			print("Week (%s) not yet completed in Season - %s", week, year)
			break
		}
		get_week(util.SEASON, util.WEEK)
	}
}

func get_week(year int, week int) {
	print(year, ":", week)
}
