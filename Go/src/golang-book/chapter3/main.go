package main

import (
       "fmt" 
       "math"
)

func main() {
        var (
	    a = 5
	    b = 34
	    c = 2342975
	)

	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World")
	fmt.Println(a, b, c)

	fmt.Println(true && true)
	fmt.Println(true || true)
	fmt.Println(!true)

	var x string
	x = "first "
	fmt.Println(x)
	x += "second"
	fmt.Println(x)

	ch := 'm'
	fmt.Println(ch)

	dogsName := "Max"
	fmt.Println("My dog's name is", dogsName)

	const importantWord string = "mouse"
	// importantWord = "mouse"

	fmt.Println(math.Pi)
}

