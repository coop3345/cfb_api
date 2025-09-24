package models

type Conferences []Conference
type Conference struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	ShortName      string `json:"shortName"`
	Abbreviation   string `json:"abbreviation"`
	Classification string `json:"classification"`
}
