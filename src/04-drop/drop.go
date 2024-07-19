package drop

import (
	"fmt"
	"time"
)

const (
	cap  = 100
	work = 2000
)

func Drop() {
	ch := make(chan string, cap)

	go func() {
		for p := range ch {
			fmt.Println("child : received signal :", p)
		}
	}()

	for i := 0; i < work; i++ {
		select {
		case ch <- "data":
			fmt.Println("parent : sent signal :", i)
		default:
			fmt.Println("parent : dropped data :", i)
		}
	}

	close(ch)
	fmt.Println("parent : shutdown signal")
	time.Sleep(time.Second)
}
