package main

import (
	"fmt"
	"time"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

func main(){
	fmt.Println("Starting NATS Server...")
	url:= "nats://localhost:4222"
	nc, err:= nats.Connect(url)
	if err  != nil{
		logrus.Fatal("Error connecting NATS Server...")
	}
	defer nc.Close()

	count:= 0
	sub, _ := nc.Subscribe("events.local", func(msg *nats.Msg){
		count++
		fmt.Printf("Msg received on sub: %v, Data: %v\n", msg.Subject, string(msg.Data))
		if msg.Reply != "" {
			msg.Respond([]byte("OK"))
		}
	})
	defer sub.Unsubscribe()

	for {
		old:= count
		time.Sleep(15*time.Second)
		if old == count {
			fmt.Println("No msg to consume..!")
			break
		}
	}
	fmt.Println("Processed %v messages\n", count)
}
