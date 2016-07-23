package main_test

import (
	"bytes"
	"testing"

	"github.com/mukulrawat1986/goweather"
	"github.com/stretchr/testify/assert"
)

// Test_E2ERun is the end to end test of the main Run function
func Test_E2ERun(t *testing.T) {

	setup()

	// Initialize an assert variable to do assert checking
	a := assert.New(t)

	// create a bytes Buffer that will serve as our source of bytes.
	// bytes Buffer also satisfies the stringReader interface
	r := bytes.NewBuffer([]byte("Dehradun\n"))

	// create another bytes Buffer that will be our sink of bytes
	w := &bytes.Buffer{}

	// call the main function
	main.Run(r, w)

	res := w.String()

	a.Contains(res, "You entered the city:")
	a.Contains(res, "Dehradun")
	a.Contains(res, "Weather Information")
	a.Contains(res, "City: Dehra Dun")
	a.Contains(res, "Weather: broken clouds")
}
