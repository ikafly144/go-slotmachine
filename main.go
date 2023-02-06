package main

import (
	"fmt"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")
	w.Resize(fyne.NewSize(480, 360))
	w.SetFixedSize(true)

	result := widget.NewRichText(&widget.TextSegment{Style: widget.RichTextStyleCodeInline, Text: ""})
	content := widget.NewButton("play", func() {
		n1, n2, n3, win := play()
		res := ""
		switch win {
		case 1:
			res = "Win!"
			w1 := a.NewWindow("Win")
			w1.SetContent(widget.NewLabel("You Won!"))
			w1.Resize(fyne.NewSize(100, 50))
			w1.SetFixedSize(true)
			w1.SetPadded(true)
			w1.RequestFocus()
			w1.Show()
		case 2:
			res = "BIG WIN 777!"
			w1 := a.NewWindow("Win")
			w1.SetContent(widget.NewLabel("You Won!"))
			w1.Resize(fyne.NewSize(100, 50))
			w1.SetFixedSize(true)
			w1.RequestFocus()
			w1.Show()
		}
		result.Segments = []widget.RichTextSegment{
			&widget.TextSegment{
				Style: widget.RichTextStyleCodeInline,
				Text:  fmt.Sprintf("|%d|%d|%d|%s\n%s", n1, n2, n3, res, result.String()),
			},
		}
		result.Refresh()
	})
	w.SetContent(container.New(layout.NewGridLayoutWithRows(2), content, container.NewScroll(result)))
	w.Show()

	a.Run()
}

func play() (int, int, int, int) {
	n1 := rand.Intn(7)
	n2 := rand.Intn(7)
	n3 := rand.Intn(7)
	var win int
	if n1 == n2 && n2 == n3 && n1 == n3 {
		win = 1
		if n1 == 7 {
			win = 2
		}
	}
	return n1, n2, n3, win
}
