package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var CityNames = []string{}

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
