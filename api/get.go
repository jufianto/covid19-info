package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

type GetAPI struct{
	GetURL string
	Client *http.Client
}

var BaseURL = "https://api.covid19api.com"

func NewGetAPI() GetAPI{
	return GetAPI{
		Client: &http.Client{},
	}
}

func (a *GetAPI) setUrl(url ...string){
	extUrl := strings.Join(url, "/")
	newUrl := fmt.Sprintf("%s/%s", BaseURL, extUrl)
	a.GetURL = newUrl
}

func(a *GetAPI) GetCountries() (countries []Countries, err error){
	now := time.Now()
	a.setUrl("countries")
	req, err := http.NewRequest(http.MethodGet, a.GetURL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := a.Client.Do(req)
	if err != nil{
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&countries)
	if err != nil{
		return nil, err
	}
	log.Println("Process get Countries", time.Since(now))
	return countries, nil
}

func (a *GetAPI) GetByCountry(countrySlug string) (covidInfo []CovidApiInfo, err error){
	now := time.Now()
	a.setUrl("country", countrySlug)
	req, err := http.NewRequest(http.MethodGet, a.GetURL, nil)
	if err != nil{
		return nil, err
	}
	resp, err := a.Client.Do(req)
	err = json.NewDecoder(resp.Body).Decode(&covidInfo)
	if err != nil{
		return nil, err
	}
	log.Println("Process get API", time.Since(now))
	return
}

func (a *GetAPI) GetCovidSummary() (covidSummary CovidApiSummary, err error){
	now := time.Now()
	a.setUrl("summary")
	req, err := http.NewRequest(http.MethodGet, a.GetURL, nil)
	if err != nil{
		return covidSummary, err
	}
	resp, err := a.Client.Do(req)
	err = json.NewDecoder(resp.Body).Decode(&covidSummary)
	if err != nil{
		return covidSummary, err
	}
	log.Println("Process get API", time.Since(now))
	return
}
