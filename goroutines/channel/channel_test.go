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
	arrInt := []int{1, 5, 10, 11, 2, 3, 4, 12, 13, 6, 7, 8, 9}

	defer close(evenChan)
	defer close(oddChan)
	for _, v := range arrInt {
		go func(v int) {
			time.Sleep(time.Second * 2)
			if v%2 == 0 {
				evenChan <- v
			} else {
				oddChan <- v
			}
		}(v)
	}
	counter := 0
	for counter < len(arrInt) {
		select {
		case d := <-evenChan:
			fmt.Println("Even : ", d)
			counter++
		case d := <-oddChan:
			fmt.Println("Odd : ", d)
			counter++
		}
	}
}

func Test_DefaultSelect(t *testing.T) {
	evenChan, oddChan := make(chan int), make(chan int)
	arrInt := []int{1, 5, 10, 11, 2, 3, 4, 12, 13, 6, 7, 8, 9}

	defer close(evenChan)
	defer close(oddChan)
	for _, v := range arrInt {
		go func(v int) {
			if v%2 == 0 {
				time.Sleep(time.Millisecond * 100)
				evenChan <- v
			} else {
				time.Sleep(time.Millisecond * 90)
				oddChan <- v
			}
		}(v)
	}

	counter := 0
	for counter < len(arrInt) {
		select {
		case d := <-evenChan:
			fmt.Println("Even : ", d)
			counter++
		case d := <-oddChan:
			fmt.Println("Odd : ", d)
			counter++
		default:
			fmt.Println("waiting for data")
		}
	}
}
