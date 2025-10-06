package seasonal

import (
	"cfbapi/conn"
	"cfbapi/util"
	"encoding/json"
	"fmt"
	"strconv"
)

type PlayerUsage struct {
	Season        int    `json:"season"`
	PlayerId      string `json:"id"`
	Name          string `json:"name"`
	Position      string `json:"position"`
	Team          string `json:"team"`
	Conference    string `json:"conference"`
	PassingDowns  float64
	StandardDowns float64
	ThirdDown     float64
	SecondDown    float64
	FirstDown     float64
	Rush          float64
	Pass          float64
	Overall       float64
}

func (p *PlayerUsage) UnmarshalJSON(data []byte) error {
	type usageStruct struct {
		PassingDowns  float64 `json:"passingDowns"`
		StandardDowns float64 `json:"standardDowns"`
		ThirdDown     float64 `json:"thirdDown"`
		SecondDown    float64 `json:"secondDown"`
		FirstDown     float64 `json:"firstDown"`
		Rush          float64 `json:"rush"`
		Pass          float64 `json:"pass"`
		Overall       float64 `json:"overall"`
	}
	type aliasPlayerUsage struct {
		Season     int         `json:"season"`
		Id         string      `json:"id"`
		Name       string      `json:"name"`
		Position   string      `json:"position"`
		Team       string      `json:"team"`
		Conference string      `json:"conference"`
		Usage      usageStruct `json:"usage"`
	}

	aux := &aliasPlayerUsage{}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	p.Season = aux.Season
	p.PlayerId = aux.Id
	p.Name = aux.Name
	p.Position = aux.Position
	p.Team = aux.Team
	p.Conference = aux.Conference
	p.PassingDowns = aux.Usage.PassingDowns
	p.StandardDowns = aux.Usage.StandardDowns
	p.ThirdDown = aux.Usage.ThirdDown
	p.SecondDown = aux.Usage.SecondDown
	p.FirstDown = aux.Usage.FirstDown
	p.Rush = aux.Usage.Rush
	p.Pass = aux.Usage.Pass
	p.Overall = aux.Usage.Overall

	return nil
}

func FetchAndInsertPlayerUsage() error {
	var pu []PlayerUsage
	query := fmt.Sprintf("player/usage?year=%v", strconv.Itoa(util.SEASON))
	conn.APICall(query, &pu)

	result := util.CONFIG.CONNECTIONS.DB.Where("season = ?", util.SEASON).Delete(&PlayerUsage{})
	if result.Error != nil {
		util.LogDBError("Delete usages failed: %v", result.Error)
	}

	fmt.Printf("Deleted %d rows\n", result.RowsAffected)
	util.LogDBError("FetchAndInsertPlayerUsage", conn.BatchInsert(util.CONFIG.CONNECTIONS.DB, pu, 100))

	return nil
}
