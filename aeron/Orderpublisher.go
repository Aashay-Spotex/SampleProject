package main

import (
	"SampleProject/core-gui"
	"bytes"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/lirm/aeron-go/aeron"
	"github.com/lirm/aeron-go/aeron/atomic"
	"github.com/lirm/aeron-go/examples"
	"log"
	"os"
	"time"
)

func main() {
	err := godotenv.Load("aeron/aeron.env")
	if err != nil {
		fmt.Println("error loading .env file: %w", err)
		return
	}
	var interrupt = make(chan os.Signal, 1)
	to := time.Duration(time.Millisecond.Nanoseconds() * *examples.ExamplesConfig.DriverTo)
	ctx := aeron.NewContext().AeronDir("c:\\\\temp\\\\aeron").MediaDriverTimeout(to)

	a, err := aeron.Connect(ctx)
	if err != nil {
		fmt.Println("Failed to connect to media driver: %s\n", err.Error())
	}
	defer a.Close()
	publication, err := a.AddPublication(os.Getenv("aeron.channel1"), int32(50))
	if err != nil {
		fmt.Println(err.Error())
	}
	defer publication.Close()
	fmt.Println("Publication found %v", publication)
	marshaller := core_gui.NewSbeGoMarshaller()
	order := &core_gui.OrderBuffer{
		AcctId:      10261,
		CreateTime:  uint64(time.Now().UnixNano()),
		InstId:      1,
		LinkId:      987654321,
		OrderQty:    1234,
		MinQty:      1000,
		Price:       1.1050,
		OrderType:   uint8(0),
		TimeInForce: uint8(2),
		BaseOrder:   uint8(1),
		RequestId:   []uint8("buy-order-001"),
	}
	for counter := 0; counter < 5; counter++ {
		var buf bytes.Buffer
		err := order.Encode(marshaller, &buf, true)
		if err != nil {
			log.Printf("Failed to encode order: %v", err)
			continue
		}
		srcBuffer := atomic.MakeBuffer(buf.Bytes())
		ret := publication.Offer(srcBuffer, 0, int32(buf.Len()), nil)
		switch ret {
		case aeron.NotConnected:
			log.Printf("%d: not connected yet", counter)
		case aeron.BackPressured:
			log.Printf("%d: back pressured", counter)
		default:
			if ret < 0 {
				log.Printf("%d: Unrecognized code: %d", counter, ret)
			} else {
				log.Printf("%d: Order sent!", counter)
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
