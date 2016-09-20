package main

import "fmt"

func addMany(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func main() {
	fmt.Println(addMany(1, 2, 3))

	addTwo := func(x, y int) int {
		return x + y
	}

	fmt.Println(addTwo(1, 1))
}
