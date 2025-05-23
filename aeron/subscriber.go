package main

import (
	"bytes"
	"fmt"

	"github.com/lirm/aeron-go/aeron"
	"github.com/lirm/aeron-go/aeron/atomic"
	"github.com/lirm/aeron-go/aeron/idlestrategy"
	"github.com/lirm/aeron-go/aeron/logbuffer"
	"github.com/lirm/aeron-go/examples"
	"time"
)

func main() {

	to := time.Duration(time.Millisecond.Nanoseconds() * *examples.ExamplesConfig.DriverTo)
	ctx := aeron.NewContext().AeronDir("c:\\\\temp\\\\aeron").MediaDriverTimeout(to)

	a, err := aeron.Connect(ctx)
	if err != nil {
		fmt.Println("Failed to connect to media driver: %s\n", err.Error())
	}
	defer a.Close()
	// channel must be specific to the individual machine
	subscription, err := a.AddSubscription("aeron:udp?endpoint=192.168.1.110:14000|ttl=30", int32(*examples.ExamplesConfig.StreamID))
	if err != nil {
		fmt.Println(err)
	}
	defer subscription.Close()
	fmt.Printf("Subscription found %v", subscription)

	tmpBuf := &bytes.Buffer{}
	counter := 1
	handler := func(buffer *atomic.Buffer, offset int32, length int32, header *logbuffer.Header) {
		bytes := buffer.GetBytesArray(offset, length)
		tmpBuf.Reset()
		buffer.WriteBytes(tmpBuf, offset, length)
		fmt.Printf("%8.d: Gots me a fragment offset:%d length: %d payload: %s (buf:%s)\n", counter, offset, length, string(bytes), string(tmpBuf.Next(int(length))))

		counter++
	}

	idleStrategy := idlestrategy.Sleeping{SleepFor: time.Millisecond}

	for {
		fragmentsRead := subscription.Poll(handler, 10)
		idleStrategy.Idle(fragmentsRead)
	}
}
