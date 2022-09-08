package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Тестовый слайдер")
	w.Resize(fyne.NewSize(500, 500))

	slider := widget.NewSlider(0.0, 10.0)

	// Установка значения слайдера по нажатию кнопки
	//entry := widget.NewEntry()
	//btn := widget.NewButton("Set Value", func() {
	//	value, _ := strconv.ParseFloat(entry.Text, 64)
	//	slider.Value = value
	//	slider.Refresh()
	//})

	// Использование OnChanged
	//label := widget.NewLabel("Slider Value...")
	//slider.OnChanged = func(value float64) {
	//	label.SetText(fmt.Sprintf("%f", value))
	//}

	// Оценка (отзыв) приложения
	title := widget.NewLabel("Оцените работу нашего приложения от 0 до 10")
	label := widget.NewLabel("Оставьте свой отзыв. Что Вам не понравилось в нашей программе?")
	feed := widget.NewLabel("Ваша оценка")
	entry := widget.NewEntry()
	entry.PlaceHolder = "Отзыв..."
	btn := widget.NewButton("Отправить отзыв", func() {
		fmt.Println(entry.Text)
	})

	label.Hide()
	entry.Hide()
	btn.Hide()

	slider.OnChanged = func(value float64) {
		feed.SetText("Ваша оценка: " + fmt.Sprintf("%f", value))

		if value < 5 {
			label.Show()
			label.SetText("Оставьте свой отзыв. Что Вам не понравилось в нашей программе?")
			entry.Show()
			btn.Show()
		} else {
			label.SetText("Спасибо за Ваш позитивный отзыв!")
			entry.Hide()
			btn.Hide()
		}
	}

	w.SetContent(
		container.NewVBox(
			title,
			slider,
			feed,
			label,
			entry,
			btn,
		),
	)
	w.Show()
	a.Run()
}
