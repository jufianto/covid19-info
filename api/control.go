package api

import (
	"log"
)

type CountriesController struct {
	Repo GetAPI
}

func (c *CountriesController) GetCovidInfoByCountries(country string) (dataCovid CovidInfo) {
	dataCovid, err := c.Repo.GetByCountry(country)
	if err != nil {
		log.Println("Error GetCovidInfoByCountries", err)
		return
	}
	return
}

func (c *CountriesController) GetCovidSummary() (dataCovid CovidInfo) {
	dataCovid, err := c.Repo.GetCovidSummary()
	if err != nil {
		log.Println("Error GetCovidSummary", err)
		return
	}
	return
}
