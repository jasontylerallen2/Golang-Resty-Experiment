package main

import (
	"testing"
	)

// TestIsValidIdentityByPhone runs tests for various numbers and names to check for valid identities
func TestIsValidIdentityByPhone(t *testing.T) {

	// TODO: fix this mock response so it can be parsed by IsValidIdentityByPhone
	//httpmock.ActivateNonDefault(resty.DefaultClient.GetClient())
	//
	//json := `
	//	{
	//	  "request": {
	//		"api_key": "7617fc790caa4f08a165450c7cd318fe",
	//		"primary.name": "Jason Allen",
	//		"primary.phone": "2089543374"
	//	  },
	//	  "primary_phone_checks": {
	//		"error": null,
	//		"warnings": [
	//
	//		],
	//		"is_valid": true,
	//		"country_code": "US",
	//		"is_commercial": null,
	//		"line_type": "Mobile",
	//		"carrier": "T-Mobile USA",
	//		"is_prepaid": null,
	//		"match_to_name": "No name found",
	//		"match_to_address": null,
	//		"subscriber": null
	//	  },
	//	  "secondary_phone_checks": null,
	//	  "primary_address_checks": null,
	//	  "secondary_address_checks": null,
	//	  "primary_email_address_checks": null,
	//	  "secondary_email_address_checks": null,
	//	  "ip_address_checks": null,
	//	  "identity_check_score": 404
	//	}
	//`
	//
	//responder := httpmock.NewBytesResponder(200, []byte(json))
	//httpmock.RegisterResponder("GET", IdentityCheckURL, responder)
	//
	//phone, countryCode, firstName, lastName := "12089543374", "US", "Jason", "Allen"
	//isValid, msg, e := IsValidIdentityByPhone(phone, countryCode, firstName, lastName)
	//if e != nil {
	//	t.Errorf("Test failed")
	//} else {
	//	fmt.Printf("Testing data -- Phone number: %s, Name: %s %s\n\tResults: isValid: %t, msg: %s", phone, firstName, lastName, isValid, msg)
	//}


	// TODO: this is actually an integration test since we're not mocking the data, which is bad. Fix! People change numbers all the time
	// and this will break
	phone, countryCode, firstName, lastName := "12089543374", "US", "Jason", "Allen"
	expectedValid := false
	expectedMsg := "No name found"
	RunTest(phone, countryCode, firstName, lastName, t, expectedValid, expectedMsg)

	phone, countryCode, firstName, lastName = "2088904722", "US", "Greg", "Allen"
	expectedValid = true
	expectedMsg = "Match"
	RunTest(phone, countryCode, firstName, lastName, t, expectedValid, expectedMsg)

	phone, countryCode, firstName, lastName = "potato", "US", "Carrot", "Cake"
	expectedValid = false
	expectedMsg = ""
	RunTest(phone, countryCode, firstName, lastName, t, expectedValid, expectedMsg)
}

//  RunTest calls IsValidIdentityByPhone and compares the returns with expected output
func RunTest(phone string, countryCode string, firstName string, lastName string, t *testing.T, expectedValid bool, expectedMsg string) {
	isValid, msg, e := IsValidIdentityByPhone(phone, countryCode, firstName, lastName)
	if e != nil {
		if t != nil {
			t.Log(e)
		}
	}
	if isValid != expectedValid || msg != expectedMsg {
		t.FailNow()
	}
}

//  BenchmarkTestIsValidIdentityByPhone calls IsValidIdentityByPhone and gives the runtime per loop after detecting
//  the number of iterations necessary to give a reliable time (b.N)
func BenchmarkTestIsValidIdentityByPhone(b *testing.B) {

	phone, countryCode, firstName, lastName := "12089543374", "US", "Jason", "Allen"
	expectedValid := false
	expectedMsg := "No name found"

	for n := 0; n < b.N; n++ {
		RunTest(phone, countryCode, firstName, lastName, nil, expectedValid, expectedMsg)
	}

}
