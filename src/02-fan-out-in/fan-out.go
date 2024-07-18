package fan_out_in

import (
	"fmt"
	"math/rand"
	"time"
)

func FanOut() {
	children := 2000
	ch := make(chan string, children)

	for i := 0; i < children; i++ {
		go func(child int) {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			ch <- "data"
			fmt.Println("child : sent signal :", child)
		}(i)
	}

	for children > 0 {
		d := <-ch
		children--

		fmt.Println(d)
		fmt.Println("parent : received signal : ", children)
	}

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------")
}
