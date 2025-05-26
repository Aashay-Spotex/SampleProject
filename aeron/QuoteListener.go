package main

import (
	core_gui "SampleProject/core-gui"
	"bytes"
	"fmt"
	"github.com/lirm/aeron-go/aeron/idlestrategy"
	"log"

	"github.com/joho/godotenv"
	"github.com/lirm/aeron-go/aeron"
	"github.com/lirm/aeron-go/aeron/atomic"
	"github.com/lirm/aeron-go/aeron/logbuffer"
	"github.com/lirm/aeron-go/examples"
	"os"
	"time"
)

func main() {
	err := godotenv.Load("aeron/aeron.env")
	if err != nil {
		fmt.Println("error loading .env file: %w", err)
		return
	}
	to := time.Duration(time.Millisecond.Nanoseconds() * *examples.ExamplesConfig.DriverTo)
	ctx := aeron.NewContext().AeronDir("c:\\\\temp\\\\aeron").MediaDriverTimeout(to)
	a, err := aeron.Connect(ctx)
	if err != nil {
		fmt.Println("Failed to connect to media driver: %s\n", err.Error())
	}
	defer a.Close()
	// channel must be specific to the individual machine
	subscription, err := a.AddSubscription(os.Getenv("aeron.channel"), int32(1))
	if err != nil {
		fmt.Println(err)
	}
	defer subscription.Close()
	fmt.Printf("Subscription found %v", subscription)
	counter := 0
	marshaller := core_gui.NewSbeGoMarshaller()
	handler := func(buffer *atomic.Buffer, offset int32, length int32, header *logbuffer.Header) {
		msgBytes := make([]byte, length)
		buffer.GetBytes(offset, msgBytes)
		reader := bytes.NewReader(msgBytes)

		var quote core_gui.QuoteListBuffer
		actingVersion := quote.SbeSchemaVersion()
		blockLength := quote.SbeBlockLength()

		err := quote.Decode(marshaller, reader, actingVersion, blockLength, true)
		if err != nil {
			log.Printf("Decode failed: %v", err)
			return
		}

		counter++
		fmt.Printf("Received quote #%d: %+v\n", counter, quote)
	}

	idleStrategy := idlestrategy.Sleeping{SleepFor: time.Millisecond}

	for {
		fragmentsRead := subscription.Poll(handler, 10)
		idleStrategy.Idle(fragmentsRead)
	}
}
func bytesReader(b []byte) *bytes.Reader {
	return bytes.NewReader(b)
}
