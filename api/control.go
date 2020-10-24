package api

import (
	"log"
	"strings"
)

type CountriesController struct{
	Repo GetAPI
}

func (c *CountriesController)FindCountriesByISO(key string) (country Countries){
	data, err := c.Repo.GetCountries()
	if err != nil{
		log.Println("Error FindCountriesByISO", err)
		return
	}
	for _, v := range data{
		if v.ISO2 == strings.ToUpper(key){
			country = v
			return
		}
	}
	return
}

func (c *CountriesController) GetCovidInfoByCountries(country string) (dataCovid CovidApiInfo) {
	data, err := c.Repo.GetByCountry(country)
	if err != nil{
		log.Println("Error GetCovidInfoByCountries", err)
		return
	}
	//get by last date update
	dataCovid = data[len(data)-1]

	// code if you want find by spesific date
	/*for _, v := range data{
		stringDate := v.Date.Format("2006-01-02")
		if stringDate == date {
			dataCovid = v
			break
		}
	}*/
	return
}

func (c *CountriesController) GetCovidSummary() CovidApiSummary{
	data, err := c.Repo.GetCovidSummary()
	if err != nil{
		log.Println("Error GetCovidSummary", err)
		return data
	}
	return data
}