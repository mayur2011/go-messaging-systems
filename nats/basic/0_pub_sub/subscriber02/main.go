package main

import (
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
	"fmt"
	"time"
)

func main(){
	fmt.Println("STARTING NATS SUBSCRIBER..!")
//	url := "nats://localhost:4222"
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		logrus.Fatal(err)
	}

	defer nc.Close()

	sub, _ := nc.Subscribe("events.local", func(msg *nats.Msg){
		fmt.Printf("Msg received on Subj: %v, Data: %v\n", msg.Subject, string(msg.Data))
	})

	time.Sleep(50 * time.Second)
	fmt.Println("DONE")
	sub.Unsubscribe()
}
