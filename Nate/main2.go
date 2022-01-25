package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	text, err := os.Open("letters.txt")

	//If there is an error print the error
	if err != nil {
		fmt.Println(err)
		return
	}

	//Create scanner to go through each line of letters.txt
	scanner := bufio.NewScanner(text)
	scanner.Split(bufio.ScanLines)

	//Create []strings to put file lines into
	emptySlice := []string{}

	// Add each line of letters.txt into a new slice
	for scanner.Scan() {
		emptySlice = append(emptySlice, scanner.Text())
	}

	// Get word that is being printed
	output := os.Args[1]
	rOutput := []rune(output)

	//Trackwhat line we are at
	line := 0

	//Create empty strings to add the words to
	ln1 := ""
	ln2 := ""
	ln3 := ""
	ln4 := ""
	ln5 := ""
	ln6 := ""
	ln7 := ""
	ln8 := ""

	//Check if "\n" is present in the string and at which index
	present := strings.Index(output, `\n`)

	// Go through Word and print each character
	for i := 0; i < len(rOutput); i++ {
		if i == present && present != -1 {
			//Print the first line 
			fmt.Println(ln1)
			fmt.Println(ln2)
			fmt.Println(ln3)
			fmt.Println(ln4)
			fmt.Println(ln5)
			fmt.Println(ln6)
			fmt.Println(ln7)
			fmt.Println(ln8)

			// Empty all strings
			ln1 = ""
			ln2 = ""
			ln3 = ""
			ln4 = ""
			ln5 = ""
			ln6 = ""
			ln7 = ""
			ln8 = ""

			// check if anmore \n
			present = present + strings.Index(output[present+1:], `\n`) + 1
			i++
		} else {

		
		for j:=0; j<len(emptySlice); j++ {
			if string(emptySlice[j]) == string(rOutput[i]) {
				line = j
				continue
			}
		}
		ln1 += emptySlice[line + 1]
		ln2 += emptySlice[line + 2]
		ln3 += emptySlice[line + 3]
		ln4 += emptySlice[line + 4]
		ln5 += emptySlice[line + 5]
		ln6 += emptySlice[line + 6]
		ln7 += emptySlice[line + 7]
		ln8 += emptySlice[line + 8]
		}
	}

	if ln1 != "" &&  ln2 != "" &&  ln3 != "" &&  ln4 != "" &&  ln5 != "" &&  ln6 != "" &&  ln7 != "" &&  ln8 != "" {
	//Print out the finalised strings
	fmt.Println(ln1)
	fmt.Println(ln2)
	fmt.Println(ln3)
	fmt.Println(ln4)
	fmt.Println(ln5)
	fmt.Println(ln6)
	fmt.Println(ln7)
	fmt.Println(ln8)
	}
	text.Close()

}

