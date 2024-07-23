package bounded_work_pooling

import (
	"fmt"
	"runtime"
	"sync"
)

func BoundedWorkPooling() {
	work := []string{"work", "work", "work", "work", 2000: "work"}

	g := runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup
	wg.Add(g)

	ch := make(chan string, g)

	for i := 0; i < g; i++ {
		go func(child int) {
			defer wg.Done()

			for wrk := range ch {
				fmt.Printf("child %d : received signal : %s\n", child, wrk)
			}

			fmt.Printf("child %d : received shutdown signal\n", child)
		}(i)
	}

	for _, wrk := range work {
		ch <- wrk
	}

	close(ch)
	wg.Wait()
}
