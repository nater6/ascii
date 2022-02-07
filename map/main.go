package main

import (
	"ascii"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("standard.txt")
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	// file to slice of string by line

	var lttrlines []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lttrlines = append(lttrlines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	// loop through file
	mapslice := ascii.Createmap(lttrlines)
	Output := os.Args[1]

	ascii.Condition(mapslice, Output)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		// Handle the error
	}

	file.Close()
}
