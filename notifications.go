// +build linux

package main

import (
	"fmt"

	"github.com/gen2brain/beeep"
)

func notify(title string, body string) {
	err := beeep.Notify(title, body, "")
	if err != nil {
		fmt.Println("Error sending notification:", err)
	}
}

func init() {
	err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	if err != nil {
		fmt.Println("Error initializing notifications:", err)
	}
}
