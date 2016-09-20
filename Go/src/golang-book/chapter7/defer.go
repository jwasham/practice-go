package main

import (
	//	"fmt"
	"os"
)

func main() {
	filename := "test.txt"

	f, _ := os.Open(filename)
	defer f.Close()

	// do a bunch of work with said file

	// f.Close() runs
}
