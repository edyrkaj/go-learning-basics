package main

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	// 1. Configure and Connect MQTT Client
	opts := mqtt.NewClientOptions().AddBroker("tcp://broker.hivemq.com:1883")
	opts.SetClientID("raspi-edge-node")

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("MQTT connection failed: %v", token.Error())
	}
	defer client.Disconnect(250)

	// 2. Setup Gobot Driver and Adaptor
	r := raspi.NewAdaptor()
	button := gpio.NewButtonDriver(r, "11") // GPIO pin 11

	// 3. Publish MQTT Message on Sensor Event
	work := func() {
		_ = button.On(gpio.Push, func(data interface{}) {
			payload := fmt.Sprintf(`{"event": "button_pressed", "timestamp": %d}`, time.Now().Unix())
			token := client.Publish("sensors/pi/button", 0, false, payload)
			token.Wait()
			fmt.Println("Published MQTT payload:", payload)
		})
	}

	robot := gobot.NewRobot("sensorBot",
		[]gobot.Connection{r},
		[]gobot.Device{button},
		work,
	)

	robot.Start()
}
