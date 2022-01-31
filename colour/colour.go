package main

import (
	"ascii"
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

var (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"
)

func init() {
	if runtime.GOOS == "windows" {
		Reset = ""
		Red = ""
		Green = ""
		Yellow = ""
		Blue = ""
		Purple = ""
		Cyan = ""
		Gray = ""
		White = ""
	}
}

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

	// find a string in a scanned file

	mapslice := ascii.Createmap(lttrlines)

	// Splits on newlines by default.

	ln1 := ""
	ln2 := ""
	ln3 := ""
	ln4 := ""
	ln5 := ""
	ln6 := ""
	ln7 := ""
	ln8 := ""

	// line := 0
	Output := os.Args[1]
	SlcOutput := []rune(Output)
	dash := strings.Index(Output, `\n`)

	// art := ascii.Condition(mapslice, Output)
	ascii.Condition(mapslice, Output)
	for i := range SlcOutput {
		if i == dash && dash >= 0 {
			if ln1 != "" && ln2 != "" && ln3 != "" && ln4 != "" && ln5 != "" && ln6 != "" && ln7 != "" && ln8 != "" {
				fmt.Println(ln1)
				fmt.Println(ln2)
				fmt.Println(ln3)
				fmt.Println(ln4)
				fmt.Println(ln5)
				fmt.Println(ln6)
				fmt.Println(ln7)
				fmt.Println(ln8)
			} else {
				fmt.Println()
			}

			ln1 = ""
			ln2 = ""
			ln3 = ""
			ln4 = ""
			ln5 = ""
			ln6 = ""
			ln7 = ""
			ln8 = ""
			dash = dash + strings.Index(Output[dash+1:], `\n`) + 1
			i++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		// Handle the error
	}

	println(White + "This is White" + Reset)
	println(Red + "This is Red" + Reset)
	println(Green + "This is Green" + Reset)
	println(Yellow + "This is Yellow" + Reset)
	println(Blue + "This is Blue" + Reset)
	println(Purple + "This is Purple" + Reset)
	println(Cyan + "This is Cyan" + Reset)
	println(Gray + "This is Gray" + Reset)

	file.Close()
}
