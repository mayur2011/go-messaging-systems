package main

import (
	"github.com/nats-io/nats.go"
	//"github.com/sirupsen/loprus"
	"fmt"
	"time"
	"math/rand"
)

var (
	rg = rand.New(rand.NewSource(time.Now().Unix()))
)

func main() {

	fmt.Println("Starting NATS Publisher...")
	url:= "nats://localhost:4222"
	nc, _:= nats.Connect(url)
	defer nc.Close()

	for i:=0; i<1e5; i++ {
		s:= fmt.Sprintf("Message: %v - Data: %v",i,rg.Intn(10000))
		nc.Publish("events.local",[]byte(s))
	}
	fmt.Println("All msgs published..!")
	time.Sleep(2*time.Second)
}
