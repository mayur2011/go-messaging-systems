package main

import (
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
	"fmt"
	"math/rand"
	"time"
)

var (
	rg = rand.New(rand.NewSource(time.Now().Unix()))
)

func main(){
	url:= "nats://localhost:4222"
	nc, err:= nats.Connect(url)
	if err != nil{
		logrus.Fatal("Error connecting NATS...")
	}

	defer nc.Close()

	i:=1
	for ; i<100; i++ {
		s:= fmt.Sprintf("message: %v, data: %v",i,rg.Intn(10000))

		// NATS Request from Publishing end to Subscribing subject
		_, err:= nc.Request("events.local",[]byte(s),2*time.Second)
		if err != nil{
			logrus.Errorf("Error while requesting to publish message %v: %v", i, err)
			break
		}
	}

	fmt.Println("Published %v messages\n",i)
}
