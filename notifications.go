// +build linux darwin

package main

import (
	"log"

	"github.com/gen2brain/beeep"
)

func notify(title string, body string) {
	err := beeep.Notify(title, body, "")
	if err != nil {
		log.Println("Error sending notification:", err)
	}
}
