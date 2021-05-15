package main

import (
	"runtime"

	"github.com/gen2brain/beeep"
	"github.com/webview/webview"
	"gopkg.in/toast.v1"
)

func sendNotification(title string, body string) {
	if runtime.GOOS == "windows" {
		notification := toast.Notification{
			AppID:   "{6D809377-6AF0-444B-8957-A3773F02200E}\\Octii\\octii-desktop-windows-x64.exe",
			Title:   title,
			Message: body,
		}
		notification.Push()
	} else {
		beeep.Notify(title, body, "assets/Icon.png")
	}
}

func main() {
	debug := false
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Octii")
	w.SetSize(1200, 800, webview.HintNone)
	w.SetSize(1200, 800, webview.HintMin)
	w.Navigate("https://octii.chat")
	w.Bind("inntronNotify", sendNotification)
	w.Run()
}
