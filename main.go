package main

import (
	"covid19api/api"
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("Start The Application")
	now := time.Now()

	key := "ID"
	controller := api.CountriesController{Repo: api.NewGetAPI()}
	covidInfo := controller.GetCovidInfoByCountries(key)
	covidSummary := controller.GetCovidSummary()
	log.Printf("%s Active Cases %d \n", covidInfo.Country, covidInfo.ActiveCase)
	log.Printf("%s Death %d \n", covidInfo.Country, covidInfo.TodayDeathsCase)
	log.Printf("Total Active Cases %d", covidSummary.ActiveCase)
	log.Printf("Total Deaths Cases %d", covidSummary.TodayDeathsCase)

	fmt.Println("Total Time ", time.Since(now))
}
