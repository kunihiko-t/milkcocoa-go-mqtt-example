package main

import (
	"bufio"
	"fmt"
	"github.com/kunihiko-t/milkcocoa-go-mqtt-example/common"
	"os"
)

func main() {

	config := common.NewConfig()
	c := common.GetClient(config)
	defer c.Disconnect(250)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	go func() {

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Input text and push enter key.")
		fmt.Printf(">")
		for scanner.Scan() {
			s := fmt.Sprintf("{\"params\":{\"text\":\"%v\"}}", scanner.Text())
			fmt.Println(s)
			p := c.Publish(config.Topic, 0, true, s)
			if p.Wait() && p.Error() != nil {
				panic(p.Error())
			}
			fmt.Printf(">")
		}
		if err := scanner.Err(); err != nil {
			panic(err.Error())
		}

	}()

	common.WaitSignal()

}
