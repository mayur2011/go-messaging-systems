package main

import (
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
	"fmt"
	"time"

	"os"
	"os/signal"
	"syscall"
)

func main() {
	logrus.Println("Starting NATS...")

	// Setup NATS connection`
	url:= "nats://localhost:4222"
	nc, err := nats.Connect(url)
	if err != nil{
		logrus.Fatal("Error connecting to NATS :%v", err)
	}
	defer nc.Close()

	// Initialize and start the NATS subscriber
	sub, err := nc.SubscribeSync("events.*")
	if err != nil{
		logrus.Fatal("Error initializing subscriber: %v", err)
	}
	defer sub.Unsubscribe()

	// Graceful shutdown
	shutdown:= make(chan os.Signal,1)
	signal.Notify(shutdown,syscall.SIGINT,syscall.SIGTERM)

	go func() {
		// Blocking point until signal comes
		sig:= <-shutdown
		fmt.Println("Ctrl+C -->")
		logrus.Println("Received %s signal, shuting down...", sig)

		// Graceful closing subscriber
		sub.Unsubscribe()
		nc.Close()

		// Wait for other resources to be get release
		time.Sleep(2*time.Second)
		os.Exit(0)
	}()

	fmt.Println("NATS Subsriber is running. Press Ctrl+C to gracefully shut down.")

	for {
		msg, _ := sub.NextMsg(5*time.Second); 
		if msg != nil {
			fmt.Printf("Msg rec on subject: %v & Data: %v\n",
			msg.Subject,
			string(msg.Data))
		}
		if msg == nil {
			logrus.Println("Waiting to consume messages\n")
		}
	}
}
