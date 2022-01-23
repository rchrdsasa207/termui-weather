package main

import (
	"fmt"
	"image"

	"github.com/gizak/termui/v3/widgets"
)

func WeatherIcon(title string, icon image.Image, x1, y1, x2, y2 int) *widgets.Image {
	img := widgets.NewImage(icon)
	img.SetRect(x1, y1, x2, y2)
	img.Title = title
	return img
}

func WeatherTemperature(w Weather, x1, y1, x2, y2 int) *widgets.Paragraph {
	wdgt := widgets.NewParagraph()
	wdgt.SetRect(x1, y1, x2, y2)
	wdgt.Title = "Temperature"
	wdgt.Text = fmt.Sprintf(`Temperature: %16v°C
Feels like: %17v°C
Min temperature: %12v°C
Max temperature: %12v°C
`, w.Temperature, w.FeelsLike, w.TemperatureMin, w.TemperatureMax)
	return wdgt
}

func WeatherExtras(w Weather, wnd Wind, vis int, x1, y1, x2, y2 int) *widgets.Paragraph {
	wdgt := widgets.NewParagraph()
	wdgt.SetRect(x1, y1, x2, y2)
	wdgt.Title = "Other info"
	wdgt.Text = fmt.Sprintf(`Pressure: %21v
Humidity: %20v%%
Visibility: %19v
Wind speed: %19v
`, w.Pressure, w.Humidity, vis, wnd.Speed)
	return wdgt
}
