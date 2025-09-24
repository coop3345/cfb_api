package seasonal

import (
	"cfbapi/conn"
	"cfbapi/util"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type Calendar []Week
type Week struct {
	Season     int       `json:"season" gorm:"primaryKey"`
	Week       int       `json:"week" gorm:"primaryKey"`
	SeasonType string    `json:"seasonType" gorm:"primaryKey;size:50"`
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
	b, _ := conn.APICall(fmt.Sprintf("calendar?year=%v", strconv.Itoa(util.SEASON)))
	var cal Calendar
	if err := json.Unmarshal(b, &cal); err != nil {
		panic(err)
	}

	InsertCalendar(cal)
	return cal
}

func InsertCalendar(c Calendar) {

}
