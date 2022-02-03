package main

import (
	"ascii"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("letters.txt")
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

	// new file to read
	fmt.Println(os.Args[1][10:])
	reverse, err := os.Open(os.Args[1][10:])
	if err != nil {
		fmt.Println("Usage: go run . [OPTION]")
		fmt.Print("EX: go run . --reverse=<fileName>")
		return
	}
	// file to slice of string by line

	var strreverse []string
	scanner2 := bufio.NewScanner(reverse)
	scanner2.Split(bufio.ScanLines)

	for scanner2.Scan() {
		strreverse = append(strreverse, scanner2.Text())
	}

	if err := scanner2.Err(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(strreverse)
	// loop through file
	mapslice := ascii.Createmap(lttrlines)

	checkascii := make([]string, 8)
	output := ""
	fmt.Println(len(strreverse))

	for i := 0; i <= len(strreverse[0])-1; i++ {

		for j, line := range strreverse {
			slice := []rune(line)
			checkascii[j] += string(slice[i])

		}
		match := true
		for k, art := range mapslice {
			count := 0
			for num, slice := range art {
				if num > 0 && slice == checkascii[count] {
					count++
					fmt.Println("yay")
				} else {
					match = false
					break
				}
			}
			if match {
				fmt.Println("yay2")
				output += string(rune(k))
				checkascii = make([]string, 8)
				break
			}
		}

	}

	// Output := os.Args[1]

	// art := ascii.Condition(mapslice, Output)
	// ascii.Condition(mapslice, Output)
	fmt.Println(output)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		// Handle the error
	}

	file.Close()
}
