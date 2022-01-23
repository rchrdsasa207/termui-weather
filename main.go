package main

import (
	"log"

	tui "github.com/gizak/termui/v3"
	"github.com/nfnt/resize"
	"github.com/oliamb/cutter"
)

func main() {
	if err := tui.Init(); err != nil {
		log.Fatal(err)
	}
	defer tui.Close()

	weather, err := CurrentWeather()
	if err != nil {
		log.Fatal(err)
	}
	weatherImg, weatherMain, err := weatherIconData(weather.Weather)
	if err == nil {
		weatherImg, err = cutter.Crop(weatherImg, cutter.Config{
			Width:  25,
			Height: 25,
			Mode:   cutter.Centered,
		})
	}
	if err != nil {
		log.Fatal(err)
	}
	weatherImg = resize.Resize(40, 25, weatherImg, resize.Lanczos2)
	if err != nil {
		log.Fatal(err)
	}

	tui.Render(
		WeatherIcon(weatherMain, weatherImg, 0, 0, 40, 19),
		WeatherTemperature(weather.Main, 41, 0, 75, 6),
		WeatherExtras(weather.Main, weather.Wind, weather.Visibility, 41, 6, 75, 12),
	)
	for e := range tui.PollEvents() {
		if e.Type == tui.KeyboardEvent {
			break
		}
	}
}
