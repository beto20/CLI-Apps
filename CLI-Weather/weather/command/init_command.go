package command

import (
	"encoding/json"
	"fmt"
	"os"

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
		NewWeather().GetWeatherCommand(arg)
	case flag == "-f":
		NewForecast().GetForecastCommand(arg)
	case flag == "-l":
		NewLocation().GetLocationsCommand(arg)
	default:
		showHelp()
	}
}

func showHelp() {
	commands := readCommandsJson()

	var com command
	var arr []command

	for _, c := range commands {
		com = command{
			Name:    c.Name,
			Full:    c.Full,
			Short:   c.Short,
			Example: c.Example,
		}
		arr = append(arr, com)
	}

  table := simpletable.New()
  table.Header = &simpletable.Header{
    Cells: []*simpletable.Cell{
      { Align: simpletable.AlignCenter, Text: "Weather -flag" },
      { Align: simpletable.AlignCenter, Text: "Description" },
      { Align: simpletable.AlignCenter, Text: "Example" }, 
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
  table.Print()
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

	input := os.Args[0:]

	if len(input) == 1 {
		fmt.Println("help")
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
		}
	}
}

func showVersion() {
  fmt.Print("current version: ", util.VERSION)
}
