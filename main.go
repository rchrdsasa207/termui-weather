package main

import (
	"encoding/json"
	"fmt"
	"image"
	_ "image/png"
	"log"
	"net/http"
	"time"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/nfnt/resize"
	"github.com/oliamb/cutter"
)

type (
	basic struct {
		Temp       float64 `json:"temp"`
		Feels_like float64 `json:"feels_like"`
		Temp_min   float64 `json:"temp_min"`
		Temp_max   float64 `json:"temp_max"`
		Pressure   float64 `json:"pressure"`
		Humidity   float64 `json:"humidity"`
	}
	wether_type struct {
		Icon string `json:"icon"`
		Type string `json:"main"`
	}
	wind struct {
		Speed float64 `json:"speed"`
	}
	weather struct {
		Basic      basic         `json:"main"`
		Type       []wether_type `json:"weather"`
		Visibility int           `json:"visibility"`
		Wind       wind          `json:"wind"`
	}
)

const (
	url = "https://fcc-weather-api.glitch.me/api/current?lat=56.948375&lon=24.108486"
)

func GetWeather() (weather, error) {
	client := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return weather{}, err
	}
	res, err := client.Do(req)
	if err != nil {
		return weather{}, err
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	dec := json.NewDecoder(res.Body)
	response := weather{}
	if err := dec.Decode(&response); err != nil {
		log.Fatal(err)
	}
	return response, nil
}

func getImage(url string) (image.Image, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	image, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, err
	}
	return image, nil
}

func main() {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()
	wether_now, err := GetWeather()
	if err != nil {
		log.Fatal(err)
	}
	icon, err := getImage(wether_now.Type[0].Icon)
	if err != nil {
		log.Fatal(err)
	}
	icon, err = cutter.Crop(icon, cutter.Config{
		Width:  25,
		Height: 25,
		Mode:   cutter.Centered,
	})
	icon = resize.Resize(40, 25, icon, resize.Lanczos2)
	if err != nil {
		log.Fatal(err)
	}
	var render []ui.Drawable
	img := widgets.NewImage(icon)
	img.SetRect(0, 0, 40, 19)
	img.Title = wether_now.Type[0].Type
	render = append(render, img)
	temperatuere := widgets.NewParagraph()
	temperatuere.SetRect(41, 0, 75, 6)
	temperatuere.Title = "Temperature"
	temperatuere.Text += fmt.Sprintf("Temperature: %16v°C\n", wether_now.Basic.Temp)
	temperatuere.Text += fmt.Sprintf("Feels like: %17v°C\n", wether_now.Basic.Feels_like)
	temperatuere.Text += fmt.Sprintf("Min temperature: %12v°C\n", wether_now.Basic.Temp_min)
	temperatuere.Text += fmt.Sprintf("Max temperature: %12v°C\n", wether_now.Basic.Temp_max)
	render = append(render, temperatuere)
	other := widgets.NewParagraph()
	other.SetRect(41, 6, 75, 12)
	other.Title = "Other info"
	other.Text += fmt.Sprintf("Pressure: %21v\n", wether_now.Basic.Pressure)
	other.Text += fmt.Sprintf("Humidity: %20v%%\n", wether_now.Basic.Humidity)
	other.Text += fmt.Sprintf("Visibility: %19v\n", wether_now.Visibility)
	other.Text += fmt.Sprintf("Wind speed: %19v\n", wether_now.Wind.Speed)
	render = append(render, other)
	ui.Render(render...)
	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}
}
