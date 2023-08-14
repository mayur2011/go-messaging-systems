package main

import (
"fmt"
"github.com/nats-io/nats.go"
)

func main(){
fmt.Println("NATS Implementation POC")
user:= "SYETEM"
pwd:= "SysPwd"
url := "nats://localhost:4222"
opts := []nats.Option{
	nats.UserInfo(user,pwd),
}

nc, err := nats.Connect(url, opts...)
if err != nil{
fmt.Println(err)
}
defer nc.Close()
fmt.Println("Connected to NATS Server on localhost-4222")
}
