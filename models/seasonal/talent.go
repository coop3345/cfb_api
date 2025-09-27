package seasonal

import (
	"cfbapi/conn"
	"cfbapi/util"
	"fmt"
	"strconv"
)

type TeamTalent []Talent
type Talent struct {
	Year   int    `json:"year" gorm:"primaryKey"`
	Team   string `json:"team" gorm:"primaryKey:size:100"`
	Talent int    `json:"talent"`
}

func FetchAndInsertTalent() error {
	var talent TeamTalent
	query := fmt.Sprintf("talent?year=%v", strconv.Itoa(util.SEASON))
	conn.APICall(query, &talent)
	util.LogDBError("FetchAndInsertTalent", util.DB.CreateInBatches(talent, 100).Error)

	return nil
}
