package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var CityNames = []string{}
var ApiRoot_OpenWeather = "http://api.openweathermap.org/data/2.5/"
var ApiKey_OpenWeather = ""

func init() {

}

// Run runs the endtoend program to find out weather and forecast
// makes all the function calls.
func Run(in stringReader, out io.Writer) {

	// get the city name
	AskForCity(in)

	// Print out the city name
	fmt.Fprintf(out, "\nYou entered the city:\n")
	fmt.Fprintf(out, CityNames[0])
}

func main() {

	// Call the run function
	Run(bufio.NewReader(os.Stdin), os.Stdout)
}
