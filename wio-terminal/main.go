package main

import (
	"machine"
	"time"
)

func main() {
	led:= machine.LED
	led.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})

	for {
		print("Hello, %v\n")
		led.Toggle()
		time.Sleep(1000 * time.Millisecond)
	}
}