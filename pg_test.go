package pg

import (
	"testing"
)

func TestRequestAccessToken(t *testing.T) {

	result, success := requestAccessToken(&Credentials{ClientID: "XXX", ClientSecret: "YYY", GrantType: "client_credentials"})

	t.Log(result)

	if success {
		t.Errorf("Result from Get access token %v. Data %v", success, result)
	}
}

func TestUssdPush(t *testing.T) {
	credentials := Credentials{ClientID: "XXX", ClientSecret: "YYYY", GrantType: "client_credentials"}
	pg := PG{Credentials: credentials}

	// Test Tigo Push
	result, success := pg.RequestUssdPush(UssdPushRequest{
		Channel:     "TIGO_PUSH",
		Amount:      1000,
		Reference:   "90210REFZ",
		Currency:    "TZS",
		CallbackURL: "url://callback",
		Description: "Test USSD push",
		Msisdn:      "+25565800000",
		CountryCode: "TZ",
	})

	t.Log(result)

	if success {
		t.Errorf("Result from USSD Push %v. Data %v", success, result)
	}

}
