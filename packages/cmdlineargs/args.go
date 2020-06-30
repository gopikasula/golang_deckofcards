package main

import (
	"os"
)

// Reads command line arguments
// and prints to output console
func main() {
	for i := 0; i < len(os.Args); i++ {
		println(os.Args[i])
	}
}

// go run args.go arg1 arg2
