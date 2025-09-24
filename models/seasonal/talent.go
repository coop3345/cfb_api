package seasonal

type Talent struct {
	Year   int    `json:"year" gorm:"primaryKey"`
	Team   string `json:"team" gorm:"primaryKey:size:100"`
	Talent int    `json:"talent"`
}
