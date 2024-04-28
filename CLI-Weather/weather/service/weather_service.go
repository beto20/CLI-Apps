package service

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/beto20/CLI-Wheather/weather/util"
)

type weather struct {
	Location struct {
		Name    string  `json:"name"`
		Region  string  `json:"region"`
		Country string  `json:"country"`
		Lat     float64 `json:"lat"`
		Lon     float64 `json:"lon"`
	} `json:"location"`
	Current struct {
		TempCelsius float64 `json:"temp_c"`
		WindSpeed   float64 `json:"wind_kph"`
		Humidity    float64 `json:"humidity"`
		Cloud       float64 `json:"cloud"`
	} `json:"current"`
	Forecast struct {
		Forecastday []struct {
			Date string `json:"date"`
			Day  struct {
				MaxTempCelsius float64 `json:"maxtemp_c"`
				MinTempCelsius float64 `json:"mintemp_c"`
				AvgTempCelsius float64 `json:"avgtemp_c"`
				MaxWindSpeed   float64 `json:"maxwind_kph"`
				AvgHumidity    float64 `json:"avghumidity"`
				Uv             float64 `json:"uv"`
			} `json:"day"`
			Astro struct {
				Sunrise   string `json:"sunrise"`
				Sunset    string `json:"sunset"`
				Moonrise  string `json:"moonrise"`
				Moonset   string `json:"moonset"`
				Moonphase string `json:"moon_phase"`
			} `json:"astro"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

type weatherServiceInterface interface {
	GetCurrentWeather(location string) weather
	GetForecast(location string) weather
}

func NewWeather() weatherServiceInterface {
	return &weather{}
}

func (wtr *weather) GetCurrentWeather(location string) weather {
	url := util.URL_BASE + "/current.json"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-RapidAPI-Key", util.API_KEY)
	req.Header.Add("X-RapidAPI-Host", util.API_HOST)

	q := req.URL.Query()
	q.Add("q", location)
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

	var w weather
	err = json.Unmarshal(body, &w)
	if err != nil {
		panic(err)
	}

	return w
}

func (wtr *weather) GetForecast(location string) weather {
	url := util.URL_BASE + "/forecast.json"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-RapidAPI-Key", util.API_KEY)
	req.Header.Add("X-RapidAPI-Host", util.API_HOST)

	q := req.URL.Query()
	q.Add("q", location)
	q.Add("days", "3")
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

	var w weather
	err = json.Unmarshal(body, &w)
	if err != nil {
		panic(err)
	}

	return w
}
