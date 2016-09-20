package main

import "fmt"

func main() {
	var a [5]int

	a[4] = 100

	fmt.Println(a)

	scores := [5]float64{ 98, 93, 77, 82, 83 }

	var total float64 = 0
	count := len(scores)
	for _, score := range scores {
		total += score
	}

	fmt.Println(total / float64(count))
}
