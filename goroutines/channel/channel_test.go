package channel

import (
	"testing"
	"time"
)

func ChannelIn(channel chan<- string)  {
	time.Sleep(2 * time.Second)
	channel <- "hello world"
}

func ChannelOut(channel <-chan string)  {
	time.Sleep(2 * time.Second)
}

func Test_ChannelInOut(t *testing.T) {
	channel := make(chan string)
	go ChannelIn(channel)
	go ChannelOut(channel)

	time.Sleep(3 * time.Second)
	close(channel)
}
