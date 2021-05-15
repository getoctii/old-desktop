//+build !windows

package main

import "github.com/gen2brain/beeep"

func sendNotification(title string, body string) {
	beeep.Notify(title, body, "assets/Icon.png")
}
