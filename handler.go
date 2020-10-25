package main

import (
	"covid19api/api"
	"covid19api/twilio"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func handlerTwilioPost(_ http.ResponseWriter, r *http.Request) {
	var err error
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

	waAPI := twilio.NewWhatsappAPI(ACCOUNT_ID, TOKEN, SERVICE_NUMBER, to)
	waAPI.Text = response.MessageBody
	waAPI.SendMessage()
}
