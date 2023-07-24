package channel

import (
	"fmt"
	"testing"
	"time"
)

func ChannelIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "hello world"
}

func ChannelOut(channel <-chan string) {
	time.Sleep(2 * time.Second)
}

func Test_ChannelInOut(t *testing.T) {
	channel := make(chan string)
	go ChannelIn(channel)
	go ChannelOut(channel)

	time.Sleep(3 * time.Second)
	close(channel)
}

func Test_SelectChannel(t *testing.T) {
	evenChan, oddChan := make(chan int), make(chan int)
	arrInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	for _, v := range arrInt {
		go func(v int) {
			if v%2 == 0 {
				evenChan <- v
			} else {
				oddChan <- v
			}
		}(v)
	}

	arrEven, arrOdd := []int{}, []int{}
	for i := 0; i < len(arrInt); i++ {
		select {
		case d := <-evenChan:
			arrEven = append(arrEven, d)
		case d := <-oddChan:
			arrOdd = append(arrOdd, d)
		}
	}
	fmt.Println(arrEven, arrOdd)
}
