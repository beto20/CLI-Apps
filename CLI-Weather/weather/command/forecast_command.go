package command

import (
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
	GetForecastCommand(arg string) ForecastCommand
}

func NewForecast() ForecastInterface {
	return &ForecastCommand{}
}

func (fc *ForecastCommand) GetForecastCommand(arg string) ForecastCommand {
	nw := service.NewWeather()
	forecasts := nw.GetForecast(arg)

	var fcDetails []struct {
		Date           string
		MaxTempCelsius float64
		MinTempCelsius float64
		AvgTempCelsius float64
		AvgHumidity    float64
		Sunrise        string
		Sunset         string
	}

	for _, fc := range forecasts.Forecast.Forecastday {
		x := struct {
			Date           string
			MaxTempCelsius float64
			MinTempCelsius float64
			AvgTempCelsius float64
			AvgHumidity    float64
			Sunrise        string
			Sunset         string
		}{
			Date:           fc.Date,
			MaxTempCelsius: fc.Day.MaxTempCelsius,
			MinTempCelsius: fc.Day.MinTempCelsius,
			AvgTempCelsius: fc.Day.AvgTempCelsius,
			AvgHumidity:    fc.Day.AvgHumidity,
			Sunrise:        fc.Astro.Sunrise,
			Sunset:         fc.Astro.Sunset,
		}

		fcDetails = append(fcDetails, x)
	}

	y := ForecastCommand{
		Name:            forecasts.Location.Name,
		Country:         forecasts.Location.Country,
		ForecastDetails: fcDetails,
	}

  return y
}

