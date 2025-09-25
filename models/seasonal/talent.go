package seasonal

import (
	"cfbapi/conn"
	"cfbapi/util"
	"encoding/json"
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

	b, _ := conn.APICall(query)
	if err := json.Unmarshal(b, &talent); err != nil {
		panic(err)
	}
	if err := util.DB.CreateInBatches(talent, 100).Error; err != nil {
		return err
	}

	return nil
}
