package main

import (
	"covid19api/api"
	"covid19api/twilio"
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

	waAPI := twilio.NewWhatsappAPI("AC80c8d0a3d34adf52d71fab74aa63522d", "44ad6f1289a5f9631a2ebb122f759604", "14155238886", "6285365861261")
	waAPI.Text = fmt.Sprintf("Total Active Cases %d", covidSummary.ActiveCase)
	waAPI.SendMessage()
}
