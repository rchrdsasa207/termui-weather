package main

import (
	"encoding/json"
	"image"
	_ "image/png"
	"net/http"
)

type (
	Weather struct {
		Temperature    float64 `json:"temp"`
		TemperatureMin float64 `json:"temp_min"`
		TemperatureMax float64 `json:"temp_max"`
		FeelsLike      float64 `json:"feels_like"`
		Pressure       float64 `json:"pressure"`
		Humidity       float64 `json:"humidity"`
	}
	WeatherData struct {
		Icon string `json:"icon"`
		Main string `json:"main"`
	}
	Wind struct {
		Speed float64 `json:"speed"`
	}
	CurrentWeatherResponse struct {
		Main       Weather       `json:"main"`
		Weather    []WeatherData `json:"weather"`
		Visibility int           `json:"visibility"`
		Wind       Wind          `json:"wind"`
	}
)

const url = "https://fcc-weather-api.glitch.me/api/current?lat=56.948375&lon=24.108486"

func CurrentWeather() (r CurrentWeatherResponse, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return r, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&r)
	return r, err
}

func getImage(url string) (image.Image, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	image, _, err := image.Decode(resp.Body)
	return image, err
}

func weatherIconData(w []WeatherData) (icon image.Image, main string, err error) {
	var url string
	for _, d := range w {
		if d.Icon != "" {
			url, main = d.Icon, d.Main
		}
	}
	if url == "" {
		return nil, main, nil
	}
	icon, err = getImage(url)
	return icon, main, err
}
