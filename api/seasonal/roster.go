package seasonal

type Rosters []Roster
type Roster struct {
	Id             string   `json:"id"`
	FirstName      string   `json:"firstName"`
	LastName       string   `json:"lastName"`
	Team           string   `json:"team"`
	Height         int      `json:"height"`
	Weight         int      `json:"weight"`
	Jersey         int      `json:"jersey"`
	Position       string   `json:"position"`
	HomeCity       string   `json:"homeCity"`
	HomeState      string   `json:"homeState"`
	HomeCountry    string   `json:"homeCountry"`
	HomeLatitude   float64  `json:"homeLatitude"`
	HomeLongitude  float64  `json:"homeLongitude"`
	HomeCountyFIPS string   `json:"homeCountyFIPS"`
	RecruitIds     []string `json:"recruitIds"`
}
