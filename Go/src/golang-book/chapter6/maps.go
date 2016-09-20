package main

import "fmt"

func main() {
	x := make(map[string]int)
	x["lat"] = 10
	x["long"] = 20
	x["junk"] = 598725

	fmt.Println(x)

	delete(x, "junk")

	fmt.Println(x)

	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}

	name, ok := elements["Un"]

	fmt.Println(name, ok)
}
