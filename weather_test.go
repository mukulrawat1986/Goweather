package main_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mukulrawat1986/goweather"
	"github.com/stretchr/testify/assert"
)

func Test_GetWeather_WithResults(t *testing.T) {

	a := assert.New(t)

	body := `{
  "coord": {
    "lon": 78.03,
    "lat": 30.32
  },
  "weather": [
    {
      "id": 500,
      "main": "Rain",
      "description": "light rain"
    }
  ],
  "main": {
    "temp": 294.85,
    "pressure": 860.71,
    "temp_min": 294.85,
    "temp_max": 294.85
  },
  "id": 1273313,
  "name": "Dehra Dun",
  "cod": 200
}`

	FakeServer(body, func() {

		weather, err := main.GetWeather("Dehradun")
		a.NoError(err)
		a.Equal(weather.COD, 200)
		a.Equal(weather.CityName, "Dehra Dun")
		a.Equal(weather.ID, 1273313)
		a.Equal(weather.Weather[0].ID, 500)
	})

}

func Test_GetWeather_WithNoResults(t *testing.T) {

	a := assert.New(t)

	body := `{
		"cod": 404
	}`

	FakeServer(body, func() {
		_, err := main.GetWeather("Dehradun")
		a.Error(err)
		a.Equal("There is no such city.", err.Error())
	})

}

func FakeServer(b string, f func()) {
	root := main.ApiRoot_OpenWeather

	ts := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, b)
	}))
	defer ts.Close()

	main.ApiRoot_OpenWeather = ts.URL

	f()

	main.ApiRoot_OpenWeather = root
}
