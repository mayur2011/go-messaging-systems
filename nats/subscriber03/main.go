package main
import (
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
	"fmt"
	"time"
)

func main(){
	fmt.Println("STARTING NATS..>")
	url:="nats://localhost:4222"
	nc, err := nats.Connect(url)
	if err!= nil{
		logrus.Fatal(err)
	}

	defer nc.Close()
	
	sub, err1 := nc.SubscribeSync("events.*")
	if err1 != nil {
		logrus.Fatal(err)
	}
	
	defer sub.Unsubscribe()
	
	for {
	if msg, _ := sub.NextMsg(5 * time.Second); msg != nil {
	fmt.Printf("msg rcd on subject: %v, data:%v\n",
		msg.Subject, string(msg.Data))
	} //else {
	//	break
	//}
	}	

	//sub.Unsubscribe()

}
