package twilio

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type WhatsappAPI struct {
	Client     *http.Client
	AccountSID string
	Token      string
	To         string
	From       string
	Text       string
}

func NewWhatsappAPI(accountSID, token, from, to string) WhatsappAPI {
	return WhatsappAPI{
		Client:     &http.Client{},
		AccountSID: accountSID,
		Token:      token,
		To:         to,
		From:       from,
	}
}

func (a *WhatsappAPI) SendMessage() {
	body := url.Values{}
	body.Add("From", fmt.Sprintf("whatsapp:+%s", a.From))
	body.Add("To", a.To)
	body.Add("Body", a.Text)
	URL := fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json", a.AccountSID)
	req, err := http.NewRequest(http.MethodPost, URL, strings.NewReader(body.Encode()))
	if err != nil {
		log.Println("ERROR WA", err)
	}
	req.SetBasicAuth(a.AccountSID, a.Token)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	_, err = a.Client.Do(req)
	if err != nil {
		log.Println("ERROR WA", err)
	}
}
