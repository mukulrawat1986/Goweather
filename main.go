package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

var CityNames = []string{}
var ApiRoot_OpenWeather = "http://api.openweathermap.org/data/2.5/"
var ApiKey_OpenWeather = ""

func init() {
	ApiKey_OpenWeather = os.Getenv("OWM_KEY")
}

// Run runs the endtoend program to find out weather and forecast
// makes all the function calls.
func Run(in stringReader, out io.Writer) {

	// get the city name
	AskForCity(in)

	// Print out the city name
	fmt.Fprintf(out, "\nYou entered the city:\n")
	fmt.Fprintf(out, "%s\n", CityNames[0])

	// Fetch the weather info
	weather, err := GetWeather(CityNames[0])

	if err != nil {
		log.Panic(err)
	}

	fmt.Fprintf(out, "\nWeather Information\n")
	fmt.Fprintf(out, "City: %s\n", weather.CityName)
	fmt.Fprintf(out, "Weather: %s\n", weather.Weather[0].Descr)
	fmt.Fprintf(out, "Temperature: %.2f\n", weather.Main.Temperature)
	fmt.Fprintf(out, "Max Temperature: %.2f\n", weather.Main.TemperatureMax)
	fmt.Fprintf(out, "Minimum Temperature: %.2f\n", weather.Main.TemperatureMin)
}

func main() {

	// Call the run function
	Run(bufio.NewReader(os.Stdin), os.Stdout)
}
