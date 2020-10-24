package main

import (
	"covid19api/api"
	"covid19api/twilio"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Start The Application")
	//now := time.Now()
	//
	//key := "ID"
	//controller := api.CountriesController{Repo: api.NewGetAPI()}
	//covidInfo := controller.GetCovidInfoByCountries(key)
	//covidSummary := controller.GetCovidSummary()
	//log.Printf("%s Active Cases %d \n", covidInfo.Country, covidInfo.ActiveCase)
	//log.Printf("%s Death %d \n", covidInfo.Country, covidInfo.TodayDeathsCase)
	//log.Printf("Total Active Cases %d", covidSummary.ActiveCase)
	//log.Printf("Total Deaths Cases %d", covidSummary.TodayDeathsCase)
	//
	//fmt.Println("Total Time ", time.Since(now))

	//waAPI := twilio.NewWhatsappAPI("AC80c8d0a3d34adf52d71fab74aa63522d", "44ad6f1289a5f9631a2ebb122f759604", "14155238886", "6285365861261")
	//waAPI.Text = fmt.Sprintf("Total Active Cases %d", covidSummary.ActiveCase)
	//waAPI.SendMessage()
	http.HandleFunc("/", handlerTwilioPost)
	log.Println("Listen and Serve on Port 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func handlerTwilioPost(_ http.ResponseWriter, r *http.Request) {
	var err error
	//body, _ := ioutil.ReadAll(r.Body)
	//fmt.Println(string(body))
	err = r.ParseForm()
	if err != nil {
		log.Println("Err", err)
	}

	response := twilio.ResponseMessages{}
	messageBody := twilio.CallbackResponseTwilio{}
	messageBody.Body = r.Form.Get("Body")
	if err != nil {
		log.Println("Err Decode", err)
	}
	arrMsgBody := strings.Split(strings.ToUpper(messageBody.Body), " ")
	if len(arrMsgBody) != 2 {
		response.MessageBody = "Wrong Parameter"
	}
	controller := api.CountriesController{Repo: api.NewGetAPI()}
	if arrMsgBody[0] == "CASES" && arrMsgBody[1] != "TOTAL" {
		covidInfo := controller.GetCovidInfoByCountries(arrMsgBody[1])
		response.MessageBody = fmt.Sprintf("%s CASES %d", covidInfo.Country, covidInfo.ActiveCase)
	} else if arrMsgBody[0] == "DEATHS" && arrMsgBody[1] != "TOTAL" {
		covidInfo := controller.GetCovidInfoByCountries(arrMsgBody[1])
		response.MessageBody = fmt.Sprintf("%s DEATHS %d", covidInfo.Country, covidInfo.TodayDeathsCase)
	} else if arrMsgBody[0] == "CASES" && arrMsgBody[1] == "TOTAL" {
		covidSummary := controller.GetCovidSummary()
		response.MessageBody = fmt.Sprintf("Total Active Cases %d", covidSummary.ActiveCase)
	} else if arrMsgBody[0] == "DEATHS" && arrMsgBody[1] == "TOTAL" {
		covidSummary := controller.GetCovidSummary()
		response.MessageBody = fmt.Sprintf("Total Deaths Cases %d", covidSummary.TodayDeathsCase)
	}

	to := r.Form.Get("From")

	waAPI := twilio.NewWhatsappAPI(
		"AC80c8d0a3d34adf52d71fab74aa63522d", "2ab74f2314ac8c78cecb8429a6946766",
		"14155238886", to)
	waAPI.Text = response.MessageBody
	waAPI.SendMessage()
}
