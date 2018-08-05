package main

import (
	"github.com/go-resty/resty"
	"github.com/ttacon/libphonenumber"
	"errors"
	"fmt"
	"log"
	)

const WhitePagesAPIKey = "7617fc790caa4f08a165450c7cd318fe"
const IdentityCheckURL = "https://proapi.whitepages.com/3.3/identity_check"

// main runs a call to IsValidIdentityByPhone with a sample phone #, country, and name
func main() {
	phone, countryCode, firstName, lastName := "12089543374", "US", "Jason", "Allen"
	isValid, msg, e := IsValidIdentityByPhone(phone, countryCode, firstName, lastName)
	if e != nil {
		log.Fatal(e)
	} else {
		fmt.Printf("Testing data -- Phone number: %s, Name: %s %s\n\tResults: isValid: %t, msg: %s", phone, firstName, lastName, isValid, msg)
	}
}

// IsValidIdentityByPhone takes a phone #, country code (see: https://github.com/ttacon/libphonenumber/blob/master/countrycodetoregionmap.go),
// and a first and last name uses the WhitePagesPro API to check if the given name matches that number
// It returns a bool indicating if the identity was validated, a msg from the API, and any error that occured
func IsValidIdentityByPhone(phone, countryCode, firstName, lastName string) (isValid bool, msg string, e error) {

	// validate phone #
	num, err := libphonenumber.Parse(phone, countryCode)
	if err != nil {
		return false, "", fmt.Errorf("Invalid phone number: %s", phone)
	}
	// conver to E164 format for best reliability using WhitePages API
	phone = libphonenumber.Format(num, libphonenumber.E164)

	resp, getErr := resty.
		R().
		SetResult(map[string]interface{}{}).
		SetQueryParams(map[string]string{
			"api_key": WhitePagesAPIKey,
			"primary.name": firstName + " " + lastName,
			"primary.phone": phone,
		}).
		Get(IdentityCheckURL)

	// TODO: we can test getting in here by using a bogus URL in the future for integration tests
	if getErr != nil {
		return false, "", getErr
	}

	result := resp.Result().(*map[string]interface{})

	// parse json for match_to_name under primary_phone_checks
	phone_checks := (*result)["primary_phone_checks"]

	// TODO: we can test gettin in both the next conditions with mock data as well
	if phone_checks == nil {
		return false, "", errors.New("Invalid index primary_phone_checks - WhitePagesPro may have changed their API")
	}
	match := phone_checks.(map[string] interface{})["match_to_name"].(string)
	if match == "" {
		return false, "", errors.New("Invalid index match_to_name - WhitePagesPro may have changed their API")
	}
	if match == "Match" {
		return true, match, nil
	} else {
		return false, match, nil
	}

}
