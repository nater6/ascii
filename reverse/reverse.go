package main

import (
	"ascii"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// ASCII ART file open
	file, err := os.Open("letters.txt")
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	// ASCII ART file to slice of string by line
	var lttrlines []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lttrlines = append(lttrlines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	// INPUT FILE open
	// // fmt.Println(os.Args[1][10:])
	reverse, err := os.Open(os.Args[1][10:])
	if err != nil {
		fmt.Println("Usage: go run . [OPTION]")
		fmt.Print("EX: go run . --reverse=<fileName>")
		return
	}

	// INPUT file to slice of string by line
	var strreverse []string
	scanner2 := bufio.NewScanner(reverse)
	scanner2.Split(bufio.ScanLines)
	for scanner2.Scan() {
		strreverse = append(strreverse, scanner2.Text())
	}
	if err := scanner2.Err(); err != nil {
		fmt.Println(err)
	}

	// create map of ascii art
	mapslice := ascii.Createmap(lttrlines)
	runewidth := []rune(strreverse[0])
	checkascii := make([]string, 8)
	output := ""

	// loop through each index of the first slice (first line) of the slices of the inpput file
	for i := 0; i < len(runewidth); i++ {
		// loop through each slice, which is each new line of the file
		for j, line := range strreverse {
			// add in the art by vertical line
			slice := []rune(line)
			checkascii[j] += string(slice[i])
		}
		for k, art := range mapslice {
			count := 0
			match := true
			for num, slice := range art {
				if num != 0 && slice == checkascii[count] {
					count++
				} else if num != 0 {
					match = false
					break
				}
			}
			if match {
				output += string(rune(k))
				checkascii = make([]string, 8)
				break
			}
		}

	}
	fmt.Println(output)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		// Handle the error
	}

	file.Close()
}
