package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 10; j++ {
	    fmt.Println(j)    
	}

	for k := 0; k < 13; k++ {
	    if k % 2 == 0 {
	       fmt.Println(k, "is even")
	    } else {
	       fmt.Println(k, "is odd")
	    }
	}
}
