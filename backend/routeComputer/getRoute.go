package routecomputer

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const baseUrl = "http://router.project-osrm.org/route/v1/walking/"

type RouteResponse struct {
	Code   string  `json:"code"`
	Routes []Route `json:"routes"`
}

type Route struct {
	Geometry string `json:"geometry"`
}

func GetRoute(currLat, currLong, targetLat, targetLong float64) (string, error) {
	// NOTE: the api takes long/lat NOT lat/long
	coords := fmt.Sprintf("%f,%f;%f,%f", currLong, currLat, targetLong, targetLat)
	params := "?overview=full"

	turl := fmt.Sprintf("%s%s%s", baseUrl, coords, params)
	finalUrl, err := url.Parse(turl)
	if err != nil {
		return "", err
	}

	resp, err := http.Get(finalUrl.String())
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var routeResponse RouteResponse
	err = json.Unmarshal(body, &routeResponse)
	if err != nil {
		return "", err
	}

	if len(routeResponse.Routes) > 0 {
		geometry := routeResponse.Routes[0].Geometry
		return geometry, nil
	} else {
		return "", fmt.Errorf("no routes found in reponse")
	}
}
