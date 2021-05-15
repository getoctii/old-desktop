//+build windows

package main

import "gopkg.in/toast.v1"

func sendNotification(title string, body string) {
	notification := toast.Notification{
		AppID:   "{6D809377-6AF0-444B-8957-A3773F02200E}\\Octii\\octii-desktop-windows-x64.exe",
		Title:   title,
		Message: body,
	}
	notification.Push()
}
