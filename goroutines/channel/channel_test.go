package channel

import (
	"fmt"
	"testing"
	"time"
)

func ChannelIn(channel chan<- string)  {
	time.Sleep(2 * time.Second)
	channel <- "hello world"
}

func ChannelOut(channel <-chan string)  {
	time.Sleep(2 * time.Second)
	fmt.Println(<-channel)
}

func Test_ChannelInOut(t *testing.T) {
	channel := make(chan string)
	go ChannelIn(channel)
	go ChannelOut(channel)

	time.Sleep(3 * time.Second)
	close(channel)
}

func proccessChannel(channel <-chan string, result chan<- string, group int)  {
	for data := range channel {
		result <- data + fmt.Sprintf(" - done - group=%d",group)
	}
}

func Test_ChannelRange(t *testing.T)  {
	channelIn := make(chan string, 5)
	result := make(chan string, 5)

	for i := 0; i < 5; i++ {
		go proccessChannel(channelIn, result, i+1)
	}
	
	for i := 0; i < 15; i++ {
		data := fmt.Sprintf("data ke %d", i+1)
		channelIn <- data
	}
	
	for i := 0; i < 10; i++ {
		fmt.Println(<-result)
	}
	close(channelIn)
	close(result)
}