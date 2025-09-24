package models

type PlayStatTypes []PlayStatType
type PlayStatType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
