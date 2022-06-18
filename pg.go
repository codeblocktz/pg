package pg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cast"
)

const (
	BASE_URL = "https://pg-next.codeblock.co.tz"
)

type Credentials struct {
	ClientID     string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
	GrantType    string `json:"grantType"`
}

type PG struct {
	Credentials
}

type UssdPushRequest struct {
	Channel     string  `json:"channel"`
	Amount      float32 `json:"amount"`
	Reference   string  `json:"reference"`
	Currency    string  `json:"currency"`
	CallbackURL string  `json:"callbackUrl"`
	Description string  `json:"description"`
	Msisdn      string  `json:"msisdn"`
	CountryCode string  `json:"countryCode"`
}

func (pg PG) RequestUssdPush(input UssdPushRequest) (map[string]interface{}, bool) {
	client := &http.Client{}
	result, success := requestAccessToken(&pg.Credentials)

	if !success {
		return result, success
	}

	token := cast.ToStringMapString(result["data"])["token"]

	body, err := json.Marshal(input)

	if err != nil {
		fmt.Println(err.Error())
	}

	req, httpErr := http.NewRequest("POST", fmt.Sprintf("%s/channel/ussd/push", BASE_URL), bytes.NewBuffer(body))

	if httpErr != nil {
		fmt.Println(httpErr.Error())
		return nil, false
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err.Error())
		return nil, false
	}
	defer res.Body.Close()

	var pushResult map[string]interface{}
	decodeErr := json.NewDecoder(res.Body).Decode(&pushResult)

	if decodeErr != nil {
		fmt.Println(decodeErr.Error())
	}

	if res.StatusCode != 201 {
		return pushResult, false
	}

	return pushResult, true
}

func requestAccessToken(c *Credentials) (map[string]interface{}, bool) {
	client := &http.Client{}

	body, err := json.Marshal(c)

	if err != nil {
		fmt.Println(err.Error())
	}

	req, httpErr := http.NewRequest("POST", fmt.Sprintf("%s/auth/oauth2/token", BASE_URL), bytes.NewBuffer(body))

	if httpErr != nil {
		fmt.Println(httpErr.Error())
		return nil, false
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err.Error())
		return nil, false
	}
	defer res.Body.Close()

	var result map[string]interface{}
	decodeErr := json.NewDecoder(res.Body).Decode(&result)

	if decodeErr != nil {
		fmt.Println(decodeErr.Error())
	}

	if res.StatusCode != 201 {
		return result, false
	}

	return result, true
}
