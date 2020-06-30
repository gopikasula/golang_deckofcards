package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	dup2()
}

func dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	for key, value := range counts {
		fmt.Printf("%s \t %d \n", key, value)
	}
}

func dup2() {
	files := os.Args[1:]
	counts := make(map[string]int)
	if len(files) != 0 {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Println(err)
			}
			countLines(f, counts)
		}
	} else {
		countLines(os.Stdin, counts)
	}
	for key, value := range counts {
		fmt.Printf("%s \t %d \n", key, value)
	}
}

// imp take away counts is ptr type
func countLines(r *os.File, counts map[string]int) {
	input := bufio.NewScanner(r)
	for input.Scan() {
		counts[input.Text()]++
	}
}
