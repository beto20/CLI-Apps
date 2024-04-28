package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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

type loc struct {
	Name    string `json:"name"`
	Region  string `json:"region"`
	Country string `json:"country"`
}

func Weather(location string) {
	url := "https://weatherapi-com.p.rapidapi.com/current.json"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "")
	req.Header.Add("X-RapidAPI-Host", "weatherapi-com.p.rapidapi.com")
	q := req.URL.Query()
	q.Add("q", location)
	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		panic("")
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

	// fmt.Println(res)
	// fmt.Println(string(body))
	fmt.Println(w.Location.Country)
	fmt.Println(w.Current.TempCelsius)
}

func Forecast(location string) {
	url := "https://weatherapi-com.p.rapidapi.com/forecast.json"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "")
	req.Header.Add("X-RapidAPI-Host", "weatherapi-com.p.rapidapi.com")
	q := req.URL.Query()
	q.Add("q", location)
	q.Add("days", "3")
	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var w weather
	err = json.Unmarshal(body, &w)
	if err != nil {
		panic(err)
	}

	// fmt.Println(string(body))

	fmt.Println(w.Forecast.Forecastday[0].Day.AvgTempCelsius)
	fmt.Println(w.Forecast.Forecastday[1].Day.AvgTempCelsius)
	fmt.Println(w.Forecast.Forecastday[2].Day.AvgTempCelsius)
	fmt.Println(w.Forecast.Forecastday[0].Astro.Sunset)
	fmt.Println(w.Forecast.Forecastday[1].Astro.Sunset)
	fmt.Println(w.Forecast.Forecastday[2].Astro.Sunset)
}

func getLocationsCoincidence(location string) {
	url := "https://weatherapi-com.p.rapidapi.com/search.json"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-RapidAPI-Key", "")
	req.Header.Add("X-RapidAPI-Host", "weatherapi-com.p.rapidapi.com")

	q := req.URL.Query()
	q.Add("q", location)

	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Print(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var data []loc
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}

	// fmt.Print(l[0].Country)
	for _, values := range data {
		fmt.Println(values.Country, "\n", values.Name, "\n")
	}
}

type command struct {
	Name         string `json:"name"`
	Full         string `json:"full"`
	Short        string `json:"short"`
	RequireArgs  bool   `json:"requireArgs"`
	QuantityArgs int64  `json:"quantityArgs"`
	Description  string `json:"description"`
	Example      string `json:"example"`
}

const (
	COMMANDS_FILE = "../data/command.json"
)

func Init() {
	var f string
	var a string
	commands := readCommandsJson()

	input := os.Args[0:]

	// fmt.Println(input)

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

func chooseCommand(flag string, arg string) {
	switch {
	case flag == "-h":
		showCommands()
	case flag == "-v":
		fmt.Println(flag, arg)
	case flag == "-w":
		Weather(arg)
	case flag == "-f":
		Forecast(arg)
	case flag == "-l":
		getLocationsCoincidence(arg)
	default:
		fmt.Println("help def")
	}
}

func showExamples() {
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

	fmt.Println(arr)
}

func showCommands() {
	commands := readCommandsJson()

	var com command
	var arr []command

	for _, c := range commands {
		com = command{
			Name:         c.Name,
			Full:         c.Full,
			Short:        c.Short,
			RequireArgs:  c.RequireArgs,
			QuantityArgs: c.QuantityArgs,
			Description:  c.Description,
			Example:      c.Example,
		}
		arr = append(arr, com)
	}

	fmt.Printf(
		"Name: %v Full: %v \nShort: %v \nRequire args: %v \nDescription: %v", arr[0].Name, arr[0].Full, arr[0].Short, arr[0].RequireArgs, arr[0].Description,
	)
}

func readCommandsJson() []command {
	file, err := os.Open(COMMANDS_FILE)
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
