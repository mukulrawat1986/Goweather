package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type WeatherData struct {
	Cordinate Coord     `json:"coord"`
	Weather   []Weather `json:"weather`
	Main      MainData  `json:"main"`
	ID        int       `json:"id"`
	CityName  string    `json:"name,omitempty"`
	COD       int       `json:"cod"`
}

type Coord struct {
	Longitude float64 `json:"lon"`
	Latitude  float64 `json:"lat"`
}

type Weather struct {
	ID    int    `json:"id"`
	Main  string `json:"main"`
	Descr string `json:"description"`
}

type MainData struct {
	Temperature    float64 `json:"temp"`
	Pressure       float64 `json:"pressure"`
	TemperatureMin float64 `json:"temp_min"`
	TemperatureMax float64 `json:"temp_max"`
}

func (m *MainData) KelvinToCelsius() {
	m.Temperature -= 273
	m.TemperatureMax -= 273
	m.TemperatureMin -= 273
}

// GetWeather gets weather data from OpenWeatherMap
func GetWeather(city string) (WeatherData, error) {

	// set up the url
	u := fmt.Sprintf("%s/weather?APPID=%s&q=%s", ApiRoot_OpenWeather, ApiKey_OpenWeather, url.QueryEscape(city))

	data := WeatherData{}

	// Make the HTTP GET request
	res, err := http.Get(u)

	if err != nil {
		return WeatherData{}, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&data)

	if err != nil {
		return WeatherData{}, err
	}

	if data.COD == 404 {
		return WeatherData{}, fmt.Errorf("There is no such city.")
	}

	data.Main.KelvinToCelsius()

	return data, nil
}
