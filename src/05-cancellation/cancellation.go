package cancellation

import (
	"context"
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func Cancellation() {
	duration := 150 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	ch := make(chan string, 1)

	go func() {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		ch <- "data"
	}()

	select {
	case d := <-ch:
		fmt.Println("work complete :", d)
	case <-ctx.Done():
		fmt.Println("work cancelled")
	}

	time.Sleep(time.Second)
	fmt.Println("----------------------------------")
}

// Exercise HackerRank Walmart
func romanToOrdinal(n string) int {
	switch n {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	default:
		return 0
	}
}

type ArrName struct {
	Name     string
	Number   int
	FullName string
}

func SortNameOrdinal(names []string) []string {
	var newArr []ArrName
	var result []string

	for _, v := range names {
		acc := 0
		firstN := 0
		nameSplit := strings.Split(v, " ")
		for k, b := range nameSplit[1] {
			ordinalStr := fmt.Sprintf("%c", b)
			ordinal := romanToOrdinal(strings.ToUpper(ordinalStr))

			if k == 1 {
				firstN = ordinal
				acc += ordinal
			} else {
				if ordinal >= firstN {
					acc += ordinal
				} else {
					acc -= ordinal
				}
			}
		}

		newName := ArrName{
			Name:     nameSplit[0],
			Number:   acc,
			FullName: v,
		}

		newArr = append(newArr, newName)
	}

	sort.Slice(newArr, func(i, j int) bool {
		return newArr[i].Number < newArr[j].Number
	})

	sort.Slice(newArr, func(i, j int) bool {
		return newArr[i].Name < newArr[j].Name
	})

	for _, v := range newArr {
		result = append(result, v.FullName)
	}

	return result
}
