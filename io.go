package main

import (
	"fmt"
	"strings"
)

type stringReader interface {
	ReadString(byte) (string, error)
}

// AskForCity asks for the current name of the city
func AskForCity(r stringReader) {
	fmt.Println("Enter the name of your city: ")
	name, _ := r.ReadString('\n')
	name = strings.TrimSpace(name)
	CityNames = append(CityNames, name)
}
