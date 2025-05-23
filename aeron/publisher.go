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
	var interrupt = make(chan os.Signal, 1)
	to := time.Duration(time.Millisecond.Nanoseconds() * *examples.ExamplesConfig.DriverTo)
	ctx := aeron.NewContext().AeronDir("c:\\\\temp\\\\aeron").MediaDriverTimeout(to)

	a, err := aeron.Connect(ctx)
	if err != nil {
		fmt.Println("Failed to connect to media driver: %s\n", err.Error())
	}
	defer a.Close()
	// channel must be specific to the individual machine
	publication, err := a.AddPublication("aeron:udp?endpoint=192.168.1.110:14000|ttl=30", int32(*examples.ExamplesConfig.StreamID))
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
