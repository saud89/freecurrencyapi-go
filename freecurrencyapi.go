package freecurrencyapi

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var apikey = ""
var client = httpClient()

const BaseUrl = "https://api.freecurrencyapi.com/v1/"

func Init(key string) {
	apikey = key
}

func httpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}
	return client
}

func apiCall(endpoint string, params map[string]string) []byte {
	if len(apikey) == 0 {
		log.Fatalf("No API key provided!")
	}

	req, err := http.NewRequest("GET", BaseUrl+endpoint, nil)
	if err != nil {
		log.Fatalf("Error Occurred. %+v", err)
	}
	q := req.URL.Query()
	q.Add("apikey", apikey)
	for key, value := range params {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	//req.Header.Set("apikey", apikey)

	log.Printf("request URL is : %v", req.URL.String())

	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request to API endpoint. %+v", err)
	}

	// Close the connection to reuse it
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Couldn't parse response body. %+v", err)
	}

	return body
}

func Status() []byte {
	return apiCall("status", nil)
}

func Currencies(params map[string]string) []byte {
	return apiCall("currencies", params)
}

func Latest(params map[string]string) []byte {
	return apiCall("latest", params)
}

func Historical(params map[string]string) []byte {
	return apiCall("historical", params)
}
