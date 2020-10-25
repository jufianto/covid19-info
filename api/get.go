package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type GetAPI struct {
	Client *http.Client
}

var BaseURL = "https://disease.sh/v3/covid-19"

func NewGetAPI() GetAPI {
	return GetAPI{
		Client: &http.Client{},
	}
}

func (a *GetAPI) GetByCountry(countrySlug string) (covidInfo CovidInfo, err error) {
	now := time.Now()
	url := fmt.Sprintf("%s/countries/%s?strict=true", BaseURL, countrySlug)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return covidInfo, err
	}
	resp, err := a.Client.Do(req)
	err = json.NewDecoder(resp.Body).Decode(&covidInfo)
	if err != nil {
		return covidInfo, err
	}
	log.Println("Process get API", time.Since(now))
	return
}

func (a *GetAPI) GetCovidSummary() (covidSummary CovidInfo, err error) {
	now := time.Now()
	url := fmt.Sprintf("%s/all", BaseURL)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return covidSummary, err
	}
	resp, err := a.Client.Do(req)
	err = json.NewDecoder(resp.Body).Decode(&covidSummary)
	if err != nil {
		return covidSummary, err
	}
	log.Println("Process get API", time.Since(now))
	return
}
