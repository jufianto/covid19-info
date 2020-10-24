package api

type CovidApiInfo struct {
	Country     string
	CountryCode string
	Confirmed   int
	Deaths      int
	Recovered   int
	Active      int
}

type CovidApiSummary struct {
	Global CovidSummaryGlobal
}

type CovidSummaryGlobal struct {
	TotalConfirmed int
	TotalDeaths    int
	TotalRecovered int
}

type Countries struct {
	Country string `json:"Country"`
	Slug    string `json:"Slug"`
	ISO2    string `json:"ISO2"`
}

type CovidInfoByCountry struct {
	Updated int `json:"updated"`
	Country string `json:"country"`
	ActiveCase string `json:"active"`
	DeathsCase string `json:"deaths"`
}