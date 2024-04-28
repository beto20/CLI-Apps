package service

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/beto20/CLI-Wheather/weather/util"
)

type location struct {
	Name    string `json:"name"`
	Region  string `json:"region"`
	Country string `json:"country"`
}

type locationServiceInterface interface {
	GetLocationsCoincidence(site string) []location
}

// func NewLocation() *location {
// 	return &location{}
// }

func NewLocation() locationServiceInterface {
	return &location{}
}

func (lct *location) GetLocationsCoincidence(site string) []location {
	url := util.URL_BASE + "/search.json"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-RapidAPI-Key", util.API_KEY)
	req.Header.Add("X-RapidAPI-Host", util.API_HOST)

	q := req.URL.Query()
	q.Add("q", site)
	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		panic(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var l []location
	err = json.Unmarshal(body, &l)
	if err != nil {
		panic(err)
	}

	return l
}
