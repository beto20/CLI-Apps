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

	fmt.Println("0: ", y.Name)
	fmt.Println("0: ", y.Country)

	fmt.Println("1a: ", y.ForecastDetails[0].Date)
	fmt.Println("1b: ", y.ForecastDetails[1].Date)
	fmt.Println("1c: ", y.ForecastDetails[2].Date)

	fmt.Println("2: ", y.ForecastDetails[1].MaxTempCelsius)
	fmt.Println("3: ", y.ForecastDetails[1].MinTempCelsius)
	fmt.Println("4: ", y.ForecastDetails[1].Sunrise)

	fmt.Println("5: ", y.ForecastDetails[0].Sunset)
	fmt.Println("5: ", y.ForecastDetails[1].Sunset)
	fmt.Println("5: ", y.ForecastDetails[2].Sunset)
}
