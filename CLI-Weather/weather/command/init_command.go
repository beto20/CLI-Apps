package command

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/alexeyco/simpletable"
	"github.com/beto20/CLI-Wheather/weather/util"
)

type command struct {
	Name         string `json:"name"`
	Full         string `json:"full"`
	Short        string `json:"short"`
	RequireArgs  bool   `json:"requireArgs"`
	QuantityArgs int64  `json:"quantityArgs"`
	Description  string `json:"description"`
	Example      string `json:"example"`
}

func chooseCommand(flag string, arg string) {
	switch {
	case flag == "-h":
		showHelp()
	case flag == "-v":
		showVersion()
	case flag == "-w":
		showWeather(arg)
	case flag == "-f":
		showForecast(arg)
	case flag == "-l":
		showLocation(arg)
	default:
		showHelp()
	}
}

func showHelp() {
	commands := readCommandsJson()

	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Weather -flag"},
			{Align: simpletable.AlignCenter, Text: "Description"},
			{Align: simpletable.AlignCenter, Text: "Example"},
		},
	}

	var cells [][]*simpletable.Cell

	for _, v := range commands {
		content := []*simpletable.Cell{
			{Text: v.Short},
			{Text: v.Description},
			{Text: v.Example},
		}

		cells = append(cells, content)
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.SetStyle(simpletable.StyleRounded)
	table.Println()
}

func showWeather(arg string) {
	weather := NewWeather().GetWeatherCommand(arg)

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Weather"},
		},
	}

	var cells [][]*simpletable.Cell

  temp := strconv.FormatFloat(weather.TempCelsius, 'f', 2, 64)

	content := []*simpletable.Cell{
		{Text: weather.Name + " - " + temp},
	}

	cells = append(cells, content)


	table.Body = &simpletable.Body{Cells: cells}

	table.SetStyle(simpletable.StyleCompact)
	table.Println()
}

func showForecast(arg string) {
	forecast := NewForecast().GetForecastCommand(arg)

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Forecast"},
		},
	}

	var cells [][]*simpletable.Cell

	for _, v := range forecast.ForecastDetails {
    min := strconv.FormatFloat(v.MinTempCelsius, 'f', 2, 64)
    max := strconv.FormatFloat(v.MaxTempCelsius, 'f', 2, 64)
		content := []*simpletable.Cell{
			{ Text: min + " - " + max },
		}

		cells = append(cells, content)
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.SetStyle(simpletable.StyleCompact)
	table.Println()
}

func showLocation(arg string) {
	locations := NewLocation().GetLocationsCommand(arg)

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Locations"},
		},
	}

	var cells [][]*simpletable.Cell

	for _, v := range locations {
		content := []*simpletable.Cell{
			{Text: v.Name + " - " + v.Country},
		}

		cells = append(cells, content)
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.SetStyle(simpletable.StyleCompact)
	table.Println()
}

func readCommandsJson() []command {
	// fmt.Print(os.Getwd())
	file, err := os.Open(util.COMMANDS_FILE)

	if err != nil {
		fmt.Println("Error open:", err)
	}
	defer file.Close()

	var command []command
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&command)
	if err != nil {
		fmt.Println("Name: ")
	}

	return command
}

func Init() {
	var f string
	var a string
	commands := readCommandsJson()
  defMssg := "Invalid flag, use -h to show available flags"
  // notFound := false

	input := os.Args[0:]

  if len(input) == 1 {
		fmt.Println(defMssg)
	}
	if len(input) == 2 {
		f = input[1]
	}
	if len(input) >= 3 {
		f = input[1]
		a = input[2]
	}

	for _, c := range commands {
		if f == c.Short {
			chooseCommand(f, a)
      break
		}
    // notFound = true
	}

  // if notFound && f != "" {
  //   fmt.Println(defMssg)
  // }
}

func showVersion() {
	fmt.Println("current version:", util.VERSION)
}
