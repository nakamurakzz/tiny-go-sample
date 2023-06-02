package main

import (
	"context"
	"machine"

	keyboard "github.com/sago35/tinygo-keyboard"
	"github.com/sago35/tinygo-keyboard/keycodes/jp"
)

func main() {
	d := keyboard.New()

	gpioPins := []machine.Pin{
		machine.KEY1,
		machine.KEY2,
		machine.KEY3,
		machine.KEY4,
	}

	for c := range gpioPins {
		gpioPins[c].Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	}

	// KeyMediaXXX will be supported starting with tinygo-0.28.
	d.AddGpioKeyboard(gpioPins, [][][]keyboard.Keycode{
		{
			{
				jp.KeyA,
				jp.KeyB,
				jp.KeyC,
				jp.KeyD,
			},
		},
	})

	d.Loop(context.Background())
}
