package xero

import (
	"bytes"
	"encoding/json"
	"github.com/charmbracelet/log"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Xero struct {
	clientId          string
	clientSecret      string
	scopes            []string
	accessToken       AccessTokenResponse
	accessTokenExpiry time.Time
}

func New(clientId string, clientSecret string, scopes []string) (x Xero, err error) {
	x = Xero{
		clientId:          clientId,
		clientSecret:      clientSecret,
		scopes:            scopes,
		accessTokenExpiry: time.Now(),
	}

	token, err := x.getAccessToken()

	//TODO: Return nicer error
	if err != nil {
		return
	}

	x.accessToken = token

	return
}

func (x *Xero) GetEmployees() {
	var employees GetEmployeesResponse
	err := x.getFromAPI("https://api.xero.com/payroll.xro/1.0/Employees", &employees)
	if err != nil {
		log.Fatalf("API request failed: %v", err)
	}
	log.Infof("Parsed response: %+v\n", employees)
}

func (x *Xero) getFromAPI(url string, responseObject interface{}) (err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	token, err := x.getAccessToken()

	if err != nil {
		return
	}

	req.Header.Add("Authorization", "Bearer "+token.AccessToken)
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Failed to read response: %v", err)
	}

	err = json.Unmarshal(body, responseObject)
	if err != nil {
		log.Fatalf("Failed to parse JSON response: %v", err)
	}

	return nil
}

func (x *Xero) getAccessToken() (token AccessTokenResponse, err error) {

	if x.accessTokenExpiry.Sub(time.Now()) > 0 {
		return x.accessToken, nil
	}

	reqURL := "https://identity.xero.com/connect/token"
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("scope", strings.Join(x.scopes, " "))

	req, _ := http.NewRequest("POST", reqURL, bytes.NewBufferString(data.Encode()))

	req.SetBasicAuth(x.clientId, x.clientSecret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Failed to read response: %v", err)
	}

	var parsedResponse AccessTokenResponse
	err = json.Unmarshal(body, &parsedResponse)
	if err != nil {
		log.Fatalf("Failed to parse JSON response: %v", err)
	}

	x.accessToken = parsedResponse
	x.accessTokenExpiry = time.Now().Add(time.Second * time.Duration(x.accessToken.ExpiresIn))

	return
}
