package weekly

type PlayerUsage struct {
	Season     int    `json:"season"`
	Id         string `json:"id"`
	Name       string `json:"name"`
	Position   string `json:"position"`
	Team       string `json:"team"`
	Conference string `json:"conference"`
	Usage      struct {
		PassingDowns  float64 `json:"passingDowns"`
		StandardDowns float64 `json:"standardDowns"`
		ThirdDown     float64 `json:"thirdDown"`
		SecondDown    float64 `json:"secondDown"`
		FirstDown     float64 `json:"firstDown"`
		Rush          float64 `json:"rush"`
		Pass          float64 `json:"pass"`
		Overall       float64 `json:"overall"`
	} `json:"usage"`
}
