package channel

import (
	"fmt"
	"testing"
	"time"
)

func Test_RaceCondition(t *testing.T) {
	var x = 0
	for i := 0; i <= 1000; i++ {
		go func() {
			for j := 0; j <= 100; j++ {
				x = x + 1
			}
		}()
	}
	time.Sleep(time.Second * 5)
	fmt.Println("Counter : ", x)
}
