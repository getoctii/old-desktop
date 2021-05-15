package main

import (
	"github.com/webview/webview"
)

func main() {
	debug := false
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Octii")
	w.SetSize(1200, 800, webview.HintNone)
	w.SetSize(800, 600, webview.HintMin)
	w.Navigate("https://octii.chat")
	w.Bind("inntronNotify", sendNotification)
	w.Run()
}
