package common

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

type Config struct {
	Topic    string
	AppId    string
	Broker   string
	Username string
	Password string
}

func NewConfig() Config {
	c := Config{}
	c.AppId = os.Getenv("MILKCOCOA_APP_ID")
	if c.AppId == "" {
		fmt.Println("Couldn't find MILKCOCOA_APP_ID on your env.")
		os.Exit(1)
	}
	c.Topic = fmt.Sprintf("%v/message/push", c.AppId)
	c.Broker = fmt.Sprintf("tcp://%v.mlkcca.com:1883", c.AppId)
	c.Username = "sdammy"
	c.Password = c.AppId
	return c
}

func GetClient(c Config) MQTT.Client {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(1000)
	opts := MQTT.NewClientOptions()
	opts.AddBroker(c.Broker)
	opts.SetUsername(c.Username)
	opts.SetPassword(c.AppId)
	opts.SetClientID(strconv.Itoa(i))
	opts.SetKeepAlive(20)
	opts.SetCleanSession(true)
	client := MQTT.NewClient(opts)
	return client
}

func WaitSignal() {
	sigs := make(chan os.Signal)
	done := make(chan bool)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		fmt.Println(sig)
		done <- true
	}()
	<-done
}
