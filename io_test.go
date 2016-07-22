package main_test

import (
	"bytes"
	"testing"

	"github.com/mukulrawat1986/goweather"
	"github.com/stretchr/testify/assert"
)

func setup() {
	main.CityNames = []string{}
}

// Test_AskForCity tests the input/ouput for AskForCity function
func Test_AskForCity(t *testing.T) {

	// Initialize a variable for assert testing
	a := assert.New(t)

	// Initialize a slice of bytes
	b := []byte("Dehradun\n")

	// Create a Buffer of bytes
	r := bytes.NewBuffer(b)

	main.AskForCity(r)

	a.Equal(len(main.CityNames), 1)
	a.Equal(main.CityNames[0], "Dehradun")
}
