package main

import (
	"fmt"
	"time"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

func main(){
	url:= "nats://localhost:4222"
	nc, _:= nats.Connect(url)
	defer nc.Close()

	// Su

}
