package shipper

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// RequestParameters struct contains parameter for SendRequest.
type RequestParameters struct {
	Ctx             context.Context
	HTTPMethod      string
	Endpoint        string
	AdditionalBody  io.Reader
	AdditionalQuery map[string]interface{}
}

// SendRequest sends request to Shipper and process the response.
func SendRequest(request *RequestParameters, response interface{}) error {
	var req, err = http.NewRequestWithContext(request.Ctx, request.HTTPMethod, request.Endpoint, request.AdditionalBody)

	if err != nil {
		log.Fatalln(err)
	}

	// Set the request headers
	req.Header.Set("User-Agent", "Shipper/")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", Conf.APIKey)

	// Set the request queries
	setQueries(req, request.AdditionalQuery)

	var client = &http.Client{}
	var resp, errRequest = client.Do(req)

	if errRequest != nil {
		log.Fatalln(errRequest)
	}

	defer resp.Body.Close()

	// Read the response body
	var body, errResponse = ioutil.ReadAll(resp.Body)

	if errResponse != nil {
		log.Fatalln(errResponse)
	}

	// Decode the response JSON into target struct
	json.Unmarshal(body, &response)

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		var failResponse map[string]map[string]string
		json.Unmarshal(body, &failResponse)

		return errors.New(failResponse["data"]["content"])
	}

	return nil
}

// setQueries sets the URL query string of endpoint.
func setQueries(req *http.Request, query map[string]interface{}) {
	q := req.URL.Query()
	q.Add("apiKey", Conf.APIKey)

	for key, val := range query {
		q.Add(key, fmt.Sprint(val))
	}

	req.URL.RawQuery = q.Encode()
}
