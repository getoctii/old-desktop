package main

import (
	"github.com/gen2brain/beeep"
	"github.com/webview/webview"
)

func sendNotification(title string, body string) {
	beeep.Notify(title, body, "assets/Icon.png")
}

func main() {
	debug := false
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Octii")
	w.SetSize(900, 600, webview.HintNone)
	w.SetSize(900, 600, webview.HintMin)
	w.Navigate("https://octii.chat")
	w.Bind("inntronNotify", sendNotification)
	w.Run()
}
