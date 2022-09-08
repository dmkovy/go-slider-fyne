package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

func main() {
	a := app.New()
	w := a.NewWindow("Тестовый слайдер")
	w.Resize(fyne.NewSize(500, 500))

	slider := widget.NewSlider(0.0, 100.0)

	entry := widget.NewEntry()
	btn := widget.NewButton("Set Value", func() {
		value, _ := strconv.ParseFloat(entry.Text, 64)
		slider.Value = value
		slider.Refresh()
	})

	w.SetContent(
		container.NewVBox(
			slider,
			entry,
			btn,
		),
	)
	w.Show()
	a.Run()
}
