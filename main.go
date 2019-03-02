package main

import (
	"fmt"
	"math/rand"
	"time"

	emitter "github.com/emitter-io/go/v2"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	aircon("aircon 1")
	aircon("aircon 2")
	aircon("aircon 3")
	sensor()
}

func aircon(name string) {
	c := connect()
	println(name, "started")

	err := c.SubscribeWithGroup("Atb634otWV8V9aNJHnITA23nEw2kHSxl", "share-demo/", "room1", func(_ *emitter.Client, msg emitter.Message) {
		println(name, "received", string(msg.Payload()))
	})
	if err != nil {
		println(err.Error())
	}
}

func sensor() {
	c := connect()
	println("sensor started")
	for i := 0; i < 30; i++ {
		if err := c.Publish("Ly18NI2YLXP0s-yFMrqd_cIq4qIpYVrb", "share-demo/", fmt.Sprintf("%d degrees", rand.Intn(40))); err != nil {
			println(err.Error())
		}

		time.Sleep(time.Second)
	}
}

func connect() *emitter.Client {
	c, err := emitter.Connect("tcp://127.0.0.1:8080", func(_ *emitter.Client, msg emitter.Message) {})
	if err != nil {
		panic(err)
	}
	return c
}
