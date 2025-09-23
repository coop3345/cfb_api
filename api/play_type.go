package api

type PlayTypes []PlayType
type PlayType struct {
	Id           int    `json:"id"`
	Text         string `json:"text"`
	Abbreviation string `json:"abbreviation"`
}
