package main

import "fmt"

func main() {
	fmt.Println("Enter a number:")

	var userNum float64
	fmt.Scanf("%f", &userNum)

	switch userNum {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Unknown Number")
	}

}
