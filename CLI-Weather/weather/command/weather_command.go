package command

import (
	"github.com/beto20/CLI-Wheather/weather/service"
)

type WeatherCommand struct {
	Name        string
	Country     string
	TempCelsius float64
	WindSpeed   float64
	Humidity    float64
	Cloud       float64
}

type WeatherInterface interface {
	GetWeatherCommand(arg string) WeatherCommand
}

func NewWeather() WeatherInterface {
	return &WeatherCommand{}
}

func (wc *WeatherCommand) GetWeatherCommand(arg string) WeatherCommand {
	nw := service.NewWeather()
	w := nw.GetCurrentWeather(arg)

	wcs := WeatherCommand{
		Name:        w.Location.Name,
		Country:     w.Location.Country,
		TempCelsius: w.Current.TempCelsius,
		WindSpeed:   w.Current.WindSpeed,
		Humidity:    w.Current.Humidity,
		Cloud:       w.Current.Cloud,
	}

	return wcs
}

