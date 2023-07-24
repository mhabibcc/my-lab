package channel

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_RaceCondition(t *testing.T) {
	var x = 0
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				x = x + 1
			}
		}()
	}
	time.Sleep(time.Second * 5)
	fmt.Println("Counter : ", x)
}

// Mutual Exclusion
func Test_Mutex(t *testing.T) {
	var x = 0
	var mutex sync.Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}
	time.Sleep(time.Second * 5)
	fmt.Println("Counter : ", x)
}
