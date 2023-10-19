package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"log"
)

func main() {
	a := app.New()
	w := a.NewWindow("HTTP Debugger Crack")
	w.Resize(fyne.NewSize(300, 200))
	w.SetFixedSize(true)

	w.SetContent(widget.NewCard("Crack HTTP Debugger", "", widget.NewButton("Crack", func() {
		log.Println("Cracking...")
		av, sn, k := crack()
		dialog.ShowInformation("Cracked", "App Version: "+av+"\nSerial Number: "+sn+"\nKey: "+k, w)
		log.Println("Done")
	})))

	w.ShowAndRun()
}
