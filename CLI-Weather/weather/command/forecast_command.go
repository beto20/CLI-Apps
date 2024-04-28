package command

import (
	"fmt"

	"github.com/beto20/CLI-Wheather/weather/service"
)

type ForecastCommand struct {
	Name            string
	Country         string
	ForecastDetails []struct {
		Date           string
		MaxTempCelsius float64
		MinTempCelsius float64
		AvgTempCelsius float64
		AvgHumidity    float64
		Sunrise        string
		Sunset         string
	}
}

type ForecastInterface interface {
	GetForecastCommand(arg string)
}

func NewForecast() ForecastInterface {
	return &ForecastCommand{}
}

func (fc *ForecastCommand) GetForecastCommand(arg string) {
	nw := service.NewWeather()
	forecasts := nw.GetForecast(arg)

	var fcdetails = ForecastCommand{}

	for _, fc := range forecasts.Forecast.Forecastday {
		x := ForecastCommand{
			ForecastDetails: []struct {
				Date           string
				MaxTempCelsius float64
				MinTempCelsius float64
				AvgTempCelsius float64
				AvgHumidity    float64
				Sunrise        string
				Sunset         string
			}{
				{Date: fc.Date},
				{MaxTempCelsius: fc.Day.MaxTempCelsius},
				{MinTempCelsius: fc.Day.MinTempCelsius},
				{AvgTempCelsius: fc.Day.AvgTempCelsius},
				{AvgHumidity: fc.Day.AvgHumidity},
				{Sunrise: fc.Astro.Sunrise},
				{Sunset: fc.Astro.Sunset},
			},
		}

		fcdetails.ForecastDetails = append(fcdetails.ForecastDetails, x.ForecastDetails...)
	}

	y := ForecastCommand{
		Name:            forecasts.Location.Name,
		Country:         forecasts.Location.Country,
		ForecastDetails: fcdetails.ForecastDetails,
	}

	fmt.Println("0: ", y.ForecastDetails[0].Sunset)
	fmt.Println("1: ", y.ForecastDetails[1].Sunset)
	fmt.Println("2: ", y.ForecastDetails[2].Sunset)
}
