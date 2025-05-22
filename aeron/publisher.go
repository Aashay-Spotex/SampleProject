package main

import (
	"fmt"
	"github.com/lirm/aeron-go/aeron"
	"github.com/lirm/aeron-go/aeron/atomic"
	"github.com/lirm/aeron-go/examples"
	"log"
	"os"
	"time"
)

func main() {
	//ctx := aeron.NewContext().MediaDriverTimeout(10 * time.Second)
	//a, err := aeron.Connect(ctx)
	//if err != nil {
	//	log.Fatalf("Failed to connect to Aeron: %v", err)
	//}
	//defer a.Close()
	//
	//publication, err := a.AddPublication("aeron:ipc", 10)
	//if err != nil {
	//	log.Fatalf("Failed to add publication: %v", err)
	//}
	//defer publication.Close()
	//
	//for !publication.IsConnected() {
	//	time.Sleep(10 * time.Millisecond)
	//}
	//
	//for i := 1; i <= 10; i++ {
	//	msg := fmt.Sprintf("Hello Aeron %d", i)
	//
	//	srcBuffer := atomic.MakeBuffer(([]byte)(msg))
	//
	//	result := publication.Offer(srcBuffer, 0, int32(len(msg)), nil)
	//	if result < 0 {
	//		fmt.Printf("Offer failed: %d\n", result)
	//	} else {
	//		fmt.Printf("Sent: %s\n", msg)
	//	}
	//	time.Sleep(1 * time.Second)
	//}

	var interrupt = make(chan os.Signal, 1)
	to := time.Duration(time.Millisecond.Nanoseconds() * *examples.ExamplesConfig.DriverTo)
	ctx := aeron.NewContext().AeronDir("c\\:\\\\temp\\\\aeron").MediaDriverTimeout(to)

	a, err := aeron.Connect(ctx)
	if err != nil {
		fmt.Println("Failed to connect to media driver: %s\n", err.Error())
	}
	defer a.Close()

	publication, err := a.AddPublication(*examples.ExamplesConfig.Channel, int32(*examples.ExamplesConfig.StreamID))
	if err != nil {
		fmt.Println(err.Error())
	}
	defer publication.Close()
	fmt.Println("Publication found %v", publication)

	for counter := 0; counter < *examples.ExamplesConfig.Messages; counter++ {
		message := fmt.Sprintf("this is a message %d", counter)
		srcBuffer := atomic.MakeBuffer(([]byte)(message))
		ret := publication.Offer(srcBuffer, 0, int32(len(message)), nil)
		switch ret {
		case aeron.NotConnected:
			log.Printf("%d: not connected yet", counter)
		case aeron.BackPressured:
			log.Printf("%d: back pressured", counter)
		default:
			if ret < 0 {
				log.Printf("%d: Unrecognized code: %d", counter, ret)
			} else {
				log.Printf("%d: success!", counter)
			}
		}

		if !publication.IsConnected() {
			log.Printf("no subscribers detected")
		}
		select {
		case <-interrupt:
			return
		default:
			time.Sleep(time.Second)
		}
	}
}
