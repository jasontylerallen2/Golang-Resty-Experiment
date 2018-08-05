# Golang-Resty-Experiment

This project contains a function IsValidIdentityByPhone that checks if a name matches the subsriber of a phone number via the WhitePagesPro API. Unit tests and a benchmark test are included.

Build:

  $go get github.com/go-resty/resty
  
  $go get github.com/ttacon/libphonenumber
  
  $go build
  
========================================================
  
Run main:

  $go run resty_experiment.go
  
========================================================
  
Run tests:

  $go test -cover
  
========================================================

Run Included Executable (runs main):

  $ ./Go_Resty_Experiment
  
========================================================

Notes:


TestIsValidIdentityByPhone is currently more of an integration test than a Unit Test, as it calls the WhitePagesPro API without mocking data. To fix this, the httpmock class from https://github.com/jarcoal/httpmock can be used; this code is commented out as I could not get the mock json to parse correctly using Resty in the allotted time. Using mock data would provide the advantage of not having to worry about if WhitePages is up or nearing it's max query limit (There are about 150 queries left now, enough to test a few more times before breaking).

results.txt contains coverage and benchmark data (running the API call and identity check in a loop)

