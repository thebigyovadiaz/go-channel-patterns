package main

import (
	"fmt"
	cancellation "go-channel-patterns/src/05-cancellation"
)

func main() {
	// Wait For Result Pattern
	// wait_result.WaitResult()

	// Fan Out - In Pattern
	// fan_out_in.FanOut()

	// Fan Out - In Pattern
	// wait_task.WaitTask()

	// Drop Pattern
	// drop.Drop()

	// Cancellation Pattern
	// cancellation.Cancellation()
	arr := []string{"Philippe IIV", "Philip XX", "Pope i", "Marcus xxVx", "Marcus x", "Pope L"}
	result := cancellation.SortNameOrdinal(arr)
	fmt.Println(result)
}
