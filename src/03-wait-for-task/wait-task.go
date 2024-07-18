package wait_task

import (
	"fmt"
	"math/rand"
	"time"
)

func WaitTask() {
	ch := make(chan string)

	go func() {
		d := <-ch
		fmt.Println("child : received signal :", d)
	}()

	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	ch <- "data"

	time.Sleep(time.Second)
}
