package testing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jcocozza/poop.map/backend/internal/model"
	"io"
	"net/http"
	"net/url"
	"testing"
)

type eteTest struct {
	method             string
	url                string
	expectedStatusCode int
	payload            any
}

// this is a list of api calls to make for the end to end tests
var tests []eteTest = []eteTest{
	// STATUS
	{method: http.MethodGet, url: "/status", expectedStatusCode: http.StatusOK},
	// POOP LOCATION GET ALL
	{method: http.MethodGet, url: "/poop-location", expectedStatusCode: http.StatusOK},
	// POOP LOCATION CREATE
	{
		method:             http.MethodPut,
		url:                "/poop-location",
		expectedStatusCode: http.StatusCreated,
		payload: model.PoopLocation{
			Name:         "test location",
			Latitude:     123.23,
			Longitude:    122345.213,
			LocationType: model.PortaPotty,
			Seasonal:     false,
			Accessible:   false,
		},
	},
	// POOP LOCATION GET SPECIFIC
	{method: http.MethodGet, url: "/poop-location/asdf", expectedStatusCode: http.StatusNotImplemented},
	// POOP LOCATION UPVOTE
	{method: http.MethodPatch, url: "/poop-location/5e8eeb82-2ed1-4b76-9b60-2d15327be74c/upvote", expectedStatusCode: http.StatusOK},
	// POOP LOCATION DOWNVOTE
	{method: http.MethodPatch, url: "/poop-location/5e8eeb82-2ed1-4b76-9b60-2d15327be74c/downvote", expectedStatusCode: http.StatusOK},
	// REVIEW CREATE FOR POOP LOCATION
	{
		method: http.MethodPut,
		url: "/poop-location/5e8eeb82-2ed1-4b76-9b60-2d15327be74c/review",
		expectedStatusCode: http.StatusCreated,
		payload: model.Review{
			PoopLocationUUID: "5e8eeb82-2ed1-4b76-9b60-2d15327be74c",
			Rating: 5,
			Comment: "foo bar baz bozo",
		},
	},
	// REVIEW GET ALL BY POOP LOCATION
	{method: http.MethodGet, url: "/poop-location/5e8eeb82-2ed1-4b76-9b60-2d15327be74c/review", expectedStatusCode: http.StatusOK},
	// REVIEW UPVOTE
	{method: http.MethodPatch, url: "/review/this-is-a-test-uuid/upvote", expectedStatusCode: http.StatusOK},
	// REVIEW DOWNVOTE
	{method: http.MethodPatch, url: "/review/this-is-a-test-uuid/downvote", expectedStatusCode: http.StatusOK},
}

func TestAPI(t *testing.T) {
	server := SetupTest()
	defer server.Close()
	parsedURL, err := url.Parse(server.URL)
	if err != nil {
		panic(err)
	}

	serverURL := parsedURL.String()
	client := &http.Client{}

	var hasError bool = false
	for _, tt := range tests {
		name := fmt.Sprintf("%s %s", tt.method, tt.url)
		t.Run(name, func(t *testing.T) {
			url := fmt.Sprintf("%s%s", serverURL, tt.url)
			var payloadBytes []byte
			var err error
			if tt.payload != nil {
				payloadBytes, err = json.Marshal(tt.payload)
				if err != nil {
					t.Fatalf("failed to marshal payload: %v", err)
					hasError = true
				}
			}

			req, err := http.NewRequest(tt.method, url, bytes.NewBuffer(payloadBytes))
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
				hasError = true
			}
			req.Header.Set("Authorization", "test_api_key")

			resp, err := client.Do(req)
			if err != nil {
				t.Fatalf("failed to perform request: %v", err)
				hasError = true
			}
			defer resp.Body.Close()

			if resp.StatusCode != tt.expectedStatusCode {
				body, _ := io.ReadAll(resp.Body)
				t.Errorf("Expected Status %d, got %d\nError: %v", tt.expectedStatusCode, resp.StatusCode, string(body))
				hasError = true
			}
		})
	}
	// only do tear down if there are no errors
	if !hasError {
		err := TearDownTest()
		if err != nil {
			panic(err)
		}
	}
}
