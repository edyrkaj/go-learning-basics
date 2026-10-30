package main

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	// 1. Initialize the Raspberry Pi adaptor
	r := raspi.NewAdaptor()

	// 2. Initialize an LED driver on GPIO pin 7 (BCM Pin 4)
	led := gpio.NewLedDriver(r, "7")

	// 3. Define the work function (loop every 1 second)
	work := func() {
		gobot.Every(1*time.Second, func() {
			led.Toggle()
		})
	}

	// 4. Assemble the robot with driver and work loop
	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{r},
		[]gobot.Device{led},
		work,
	)

	// 5. Start the robot execution loop
	robot.Start()
}
