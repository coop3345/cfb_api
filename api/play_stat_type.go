package api

type PlayStatTypes []PlayStatType
type PlayStatType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
