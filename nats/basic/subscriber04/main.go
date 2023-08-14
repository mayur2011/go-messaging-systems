package main

import (
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
	"fmt"
	"time"
)

func main(){
	url:= "nats://localhost:4222"
	nc, err:= nats.Connect(url)
	if err != nil {
		logrus.Fatal(err)
	}

	defer nc.Close()

	sub, err:= nc.SubscribeSync("events.local")
	if err != nil {
		logrus.Fatal(err)
	}

	defer sub.Unsubscribe()

	for {
		if msg, _:= sub.NextMsg(5*time.Second); msg !=nil {
			fmt.Printf("Msg received on Subject %v - Data: %v\n", msg.Subject, string(msg.Data))
		}
	}
}
