package api

type CovidInfo struct {
	Updated         int    `json:"updated"`
	Country         string `json:"country"`
	ActiveCase      int    `json:"active"`
	TodayDeathsCase int    `json:"todayDeaths"`
}
