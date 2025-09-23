package seasonal

import (
	"cfbapi/api"
	"cfbapi/util"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type Calendar []Week
type Week struct {
	Season     int       `json:"season"`
	Week       int       `json:"week"`
	SeasonType string    `json:"seasonType"`
	StartDate  time.Time `json:"startDate"`
	EndDate    time.Time `json:"endDate"`
}

func (w *Week) UnmarshalJSON(data []byte) error {
	type Alias Week // Prevent recursion
	aux := &struct {
		StartDate string `json:"startDate"`
		EndDate   string `json:"endDate"`
		*Alias
	}{
		Alias: (*Alias)(w),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	var err error
	w.StartDate, err = time.Parse(time.RFC3339, aux.StartDate)
	if err != nil {
		return fmt.Errorf("invalid startDate: %w", err)
	}

	w.EndDate, err = time.Parse(time.RFC3339, aux.EndDate)
	if err != nil {
		return fmt.Errorf("invalid endDate: %w", err)
	}

	return nil
}

func GetCalendar() Calendar {
	b, _ := api.APICall(fmt.Sprintf("calendar?year=%s", strconv.Itoa(util.SEASON)))
	var cal Calendar
	if err := json.Unmarshal(b, &cal); err != nil {
		panic(err)
	}

	InsertCalendar(cal)
	return cal
}

func InsertCalendar(c Calendar) {

}
