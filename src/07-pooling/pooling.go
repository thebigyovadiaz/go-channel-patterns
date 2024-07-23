package pooling

import (
	"fmt"
	"runtime"
	"time"
)

const work = 100

func Pooling() {
	ch := make(chan string)

	g := runtime.GOMAXPROCS(0)

	for c := 0; c < g; c++ {
		go func(child int) {
			for d := range ch {
				fmt.Printf("child %d : received signal : %s\n", child+1, d)
			}

			fmt.Printf("child %d : received shutdown signal :\n", child)
		}(c)
	}

	for w := 0; w < work; w++ {
		data := fmt.Sprintf("data n %d", w+1)
		ch <- data
		fmt.Println("parent : sent signal :", w)
	}

	close(ch)

	fmt.Println("parent : sent shutdown signal")
	time.Sleep(time.Second)
}
