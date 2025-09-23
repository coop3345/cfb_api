package seasonal

type Teams []Team
type Team struct {
	Id             int      `json:"id"`
	School         string   `json:"school"`
	Mascot         string   `json:"mascot"`
	Abbreviation   string   `json:"abbreviation"`
	AlternateNames []string `json:"alternateNames"`
	Conference     string   `json:"conference"`
	Division       string   `json:"division"`
	Classification string   `json:"classification"`
	Color          string   `json:"color"`
	AlternateColor string   `json:"alternateColor"`
	Logos          []string `json:"logos"`
	Twitter        string   `json:"twitter"`
	Location       struct {
		Id               int     `json:"id"`
		Name             string  `json:"name"`
		City             string  `json:"city"`
		State            string  `json:"state"`
		Zip              string  `json:"zip"`
		CountryCode      string  `json:"countryCode"`
		Timezone         string  `json:"timezone"`
		Latitude         float64 `json:"latitude"`
		Longitude        float64 `json:"longitude"`
		Elevation        string  `json:"elevation"`
		Capacity         int     `json:"capacity"`
		ConstructionYear int     `json:"constructionYear"`
		Grass            bool    `json:"grass"`
		Dome             bool    `json:"dome"`
	} `json:"location"`
}
