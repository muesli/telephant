package main

import "github.com/gen2brain/beeep"

func notify(title string, body string) {
	err := beeep.Notify(title, body, "")
	if err != nil {
		panic(err)
	}
}

func init() {
	err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	if err != nil {
		panic(err)
	}
}
