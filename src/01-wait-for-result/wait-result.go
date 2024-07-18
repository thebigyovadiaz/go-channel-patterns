package wait_result

import (
	"fmt"
	"math/rand"
	"time"
)

func WaitResult() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- "data sent for child"
		fmt.Println("child: sent signal")
	}()

	d := <-ch
	fmt.Println("parent : received", d)
	time.Sleep(time.Second)
	fmt.Println("-------------------------------------")
}
