package main

import (
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/kunihiko-t/milkcocoa-go-mqtt-example/common"
)

type Payload struct {
	ID     string `json:"id"`
	Ts     int64  `json:"ts"`
	Params struct {
		Text string `json:"text"`
	} `json:"params"`
}

func main() {

	config := common.NewConfig()
	c := common.GetClient(config)
	defer c.Disconnect(250)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	go func() {

		subToken := c.Subscribe(config.Topic, 0, func(client MQTT.Client, msg MQTT.Message) {
			fmt.Println("Topic=", msg.Topic(), "Payload=", string(msg.Payload()))
			payload := &Payload{}
			if err := json.Unmarshal(msg.Payload(), payload); err != nil {
				panic(err)
			}
			fmt.Printf("Text : %v\n", payload.Params.Text)
		})

		if subToken.Wait() && subToken.Error() != nil {
			panic(subToken.Error())
		}

	}()

	common.WaitSignal()
}
