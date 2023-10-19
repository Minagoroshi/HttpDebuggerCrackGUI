package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"log"
)

func main() {
	a := app.New()
	w := a.NewWindow("HTTPDebugger Crack")

	w.SetContent(widget.NewCard("HTTPDebugger Crack", "Crack HTTPDebugger", widget.NewButton("Crack", func() {
		log.Println("Cracking...")
		crack()
		log.Println("Done")
	})))

	w.ShowAndRun()

}
