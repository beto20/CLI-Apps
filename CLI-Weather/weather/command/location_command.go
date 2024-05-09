package command

import (
	"github.com/beto20/CLI-Wheather/weather/service"
)

type LocationCommand struct {
	Name    string
	Country string
}

type LocationInterface interface {
	GetLocationsCommand(arg string) []LocationCommand
}

func NewLocation() LocationInterface {
	return &LocationCommand{}
}

func (lc *LocationCommand) GetLocationsCommand(arg string) []LocationCommand {
	nl := service.NewLocation()
	locations := nl.GetLocationsCoincidence(arg)

	var arr []LocationCommand

	for _, l := range locations {
		x := LocationCommand{
			Name:    l.Name,
			Country: l.Country,
		}

		arr = append(arr, x)
	}

	return arr
}
