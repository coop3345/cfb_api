package seasonal

import (
	"cfbapi/conn"
	"cfbapi/util"
	"fmt"
	"strconv"
)

type TeamTalent []Talent
type Talent struct {
	Year   int     `json:"year" gorm:"primaryKey"`
	Team   string  `json:"team" gorm:"primaryKey:size:100"`
	Talent float64 `json:"talent"`
}

func FetchAndInsertTalent() error {
	var talent TeamTalent
	query := fmt.Sprintf("talent?year=%v", strconv.Itoa(util.SEASON))
	conn.APICall(query, &talent)
	util.LogDBError("FetchAndInsertTalent", conn.BatchInsert(util.CONFIG.CONNECTIONS.DB, talent, 1))

	return nil
}
