package main

import (
	"covid19api/api"
	"fmt"
	"log"
	"time"
)
func main(){
	fmt.Println("Start The Application")
	now := time.Now()

	key := "ID"
	controller := api.CountriesController{Repo: api.NewGetAPI()}
	getCountry := controller.FindCountriesByISO(key)
	covidInfo :=  controller.GetCovidInfoByCountries(getCountry.Slug)
	covidSummary := controller.GetCovidSummary()
	totalActiveCase := covidSummary.Global.TotalConfirmed - covidSummary.Global.TotalDeaths - covidSummary.Global.TotalRecovered
	log.Printf("%s Active Cases %d \n", getCountry.Country, covidInfo.Active)
	log.Printf("%s Death %d \n", getCountry.Country, covidInfo.Deaths)
	log.Printf("Total Active Cases %d", totalActiveCase)
	log.Printf("Total Deaths Cases %d", covidSummary.Global.TotalDeaths)


	fmt.Println("Total Time ", time.Since(now))
}
