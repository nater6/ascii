package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	var file *os.File
	var err error
	if os.Args[2] == "standard" {
		file, err = os.Open("letters.txt")
		if err != nil {
			log.Fatalf("ERROR: %s", err)
		}
	} else if os.Args[2] == "shadow" {
		file, err = os.Open("shadow.txt")
		if err != nil {
			log.Fatalf("ERROR: %s", err)
		}
	}

	//Create file to store result
	name := os.Args[3]
	n := strings.Index(name, "--output=")

	if n == 0 {
		name = name[9:]
	} else {
		fmt.Println("Usage: go run . [STRING] [BANNER] [OPTION]")
		fmt.Println("EX: go run . something standard --output=<fileName.txt>")
		return
	}

	// Make new result file
	// er2 := os.WriteFile(name, nil, 0644)
	// if er2 != nil {
	// 	log.Fatal(er2.Error())
	// }

	resultFile, er2 := os.Create(name)
	if er2 != nil {
		log.Fatal(er2.Error())
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

	// Splits on newlines by default.
	ln1 := ""
	ln2 := ""
	ln3 := ""
	ln4 := ""
	ln5 := ""
	ln6 := ""
	ln7 := ""
	ln8 := ""

	line := 0
	Output := os.Args[1]
	SlcOutput := []rune(Output)
	dash := strings.Index(Output, `\n`)

	for i := 0; i < len(SlcOutput); i++ {
		if i == dash && dash >= 0 {
			if ln1 != "" && ln2 != "" && ln3 != "" && ln4 != "" && ln5 != "" && ln6 != "" && ln7 != "" && ln8 != "" {
				_, err = io.WriteString(resultFile, ln1+"\n")
				_, err = io.WriteString(resultFile, ln2+"\n")
				_, err = io.WriteString(resultFile, ln3+"\n")
				_, err = io.WriteString(resultFile, ln4+"\n")
				_, err = io.WriteString(resultFile, ln5+"\n")
				_, err = io.WriteString(resultFile, ln6+"\n")
				_, err = io.WriteString(resultFile, ln7+"\n")
				_, err = io.WriteString(resultFile, ln8+"\n")
			} else {
				_, err = io.WriteString(resultFile, "\n")
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
		} else {
			for j, s := range lttrlines {
				line = j
				if len(s) > 0 && s == string(SlcOutput[i]) {
					break
				}

			}

			for i := line + 1; i <= line+8; i++ {
				if i == line+1 {
					ln1 = ln1 + lttrlines[i]
				}
				if i == line+2 {
					ln2 = ln2 + lttrlines[i]
				}
				if i == line+3 {
					ln3 = ln3 + lttrlines[i]
				}
				if i == line+4 {
					ln4 = ln4 + lttrlines[i]
				}
				if i == line+5 {
					ln5 = ln5 + lttrlines[i]
				}
				if i == line+6 {
					ln6 = ln6 + lttrlines[i]
				}
				if i == line+7 {
					ln7 = ln7 + lttrlines[i]
				}
				if i == line+8 {
					ln8 = ln8 + lttrlines[i]
				}

			}
			if i == len(Output)-1 {
				_, _ = io.WriteString(resultFile, ln1+"\n")
				_, err = io.WriteString(resultFile, ln2+"\n")
				_, err = io.WriteString(resultFile, ln3+"\n")
				_, err = io.WriteString(resultFile, ln4+"\n")
				_, err = io.WriteString(resultFile, ln5+"\n")
				_, err = io.WriteString(resultFile, ln6+"\n")
				_, err = io.WriteString(resultFile, ln7+"\n")
				_, err = io.WriteString(resultFile, ln8)
				ln1 = ""
				ln2 = ""
				ln3 = ""
				ln4 = ""
				ln5 = ""
				ln6 = ""
				ln7 = ""
				ln8 = ""
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		// Handle the error
	}

	file.Close()
}
