package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	//Create the map and put in all values
	cMap := make(map[string]string)
	cMap["reset"] = "\033[0m"
	cMap["red"] = "\033[31m"
	cMap["green"] = "\033[32m"
	cMap["yellow"] = "\033[33m"
	cMap["blue"] = "\033[34m"
	cMap["purple"] = "\033[35m"
	cMap["cyan"] = "\033[36m"
	cMap["gray"] = "\033[37m"
	cMap["white"] = "\033[97m"

	//Take the collor from the terminal
	color := os.Args[2]
	color = color[8:]
	fmt.Println(color)

	//Create error msg for incorrect format
	if color != "red" && color != "green" && color != "yellow" && color != "blue" && color != "purple" && color != "cyan" && color != "gray" && color != "white" {
		fmt.Println("Usage: go run . [STRING] [OPTION]")
		fmt.Println("EX: go run . something --color=<color>")
	}

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

	// Splits on newlines by default.
	ln1 := []string{}
	ln2 := []string{}
	ln3 := []string{}
	ln4 := []string{}
	ln5 := []string{}
	ln6 := []string{}
	ln7 := []string{}
	ln8 := []string{}

	line := 0
	Output := os.Args[1]
	SlcOutput := []rune(Output)
	dash := strings.Index(Output, `\n`)

	//If every letter in being colored
	if len(os.Args) == 3 {

		for i := 0; i < len(SlcOutput); i++ {
			if i == dash && dash >= 0 {
				if len(ln1) != 0 && len(ln2) != 0 && len(ln3) != 0 && len(ln4) != 0 && len(ln5) != 0 && len(ln6) != 0 && len(ln7) != 0 && len(ln8) != 0 {
				
					for j := 0; j < len(ln1); j++ {
						if j == len(ln1)-1 {
							fmt.Println(string(cMap[color]) + ln1[j] + cMap["reset"])
							continue
						}
						fmt.Print(string(cMap[color]) + ln1[j] + cMap["reset"])
					}

					for j := 0; j < len(ln2); j++ {
						if j == len(ln2)-1 {
							fmt.Println(string(cMap[color]) + ln2[j] + cMap["reset"])
							continue
						}
						fmt.Print(string(cMap[color]) + ln2[j] + cMap["reset"])
					}

					for j := 0; j < len(ln3); j++ {
						if j == len(ln3)-1 {
							fmt.Println(string(cMap[color]) + ln3[j] + cMap["reset"])
							continue
						}
						fmt.Print(string(cMap[color]) + ln3[j] + cMap["reset"])
					}

					for j := 0; j < len(ln4); j++ {
						if j == len(ln4)-1 {
							fmt.Println(string(cMap[color]) + ln4[j] + cMap["reset"])
							continue
						}
						fmt.Print(string(cMap[color]) + ln4[j] + cMap["reset"])
					}

					for j := 0; j < len(ln5); j++ {
						if j == len(ln5)-1 {
							fmt.Println(string(cMap[color]) + ln5[j] + cMap["reset"])
							continue
						}
						fmt.Print(string(cMap[color]) + ln5[j] + cMap["reset"])
					}

					for j := 0; j < len(ln6); j++ {
						if j == len(ln6)-1 {
							fmt.Println(string(cMap[color]) + ln6[j] + cMap["reset"])
							continue
						}
						fmt.Print(string(cMap[color]) + ln6[j] + cMap["reset"])
					}

					for j := 0; j < len(ln7); j++ {
						if j == len(ln7)-1 {
							fmt.Println(string(cMap[color]) + ln7[j] + cMap["reset"])
							continue
						}
						fmt.Print(string(cMap[color]) + ln7[j] + cMap["reset"])
					}

					for j := 0; j < len(ln8); j++ {
						if j == len(ln8)-1 {
							fmt.Println(string(cMap[color]) + ln8[j] + cMap["reset"])
							continue
						}
						fmt.Print(string(cMap[color]) + ln8[j] + cMap["reset"])
					}

				} else {
					fmt.Println()
				}

				ln1 = []string{}
				ln2 = []string{}
				ln3 = []string{}
				ln4 = []string{}
				ln5 = []string{}
				ln6 = []string{}
				ln7 = []string{}
				ln8 = []string{}
				dash = dash + strings.Index(Output[dash+1:], `\n`) + 1
				i++
			} else {
				for j, s := range lttrlines {
					line = j
					if len(s) > 0 && s == string(SlcOutput[i]) {
						break
					}

				}

				for k := line + 1; k <= line+8; k++ {
					if k == line+1 {
						ln1 = append(ln1, lttrlines[k])
					}
					if k == line+2 {
						ln2 = append(ln2, lttrlines[k])
					}
					if k == line+3 {
						ln3 = append(ln3, lttrlines[k])
					}
					if k == line+4 {
						ln4 = append(ln4, lttrlines[k])
					}
					if k == line+5 {
						ln5 = append(ln5, lttrlines[k])
					}
					if k == line+6 {
						ln6 = append(ln6, lttrlines[k])
					}
					if k == line+7 {
						ln7 = append(ln7, lttrlines[k])
					}
					if k == line+8 {
						ln8 = append(ln8, lttrlines[k])
					}

				}
				// if i == len(Output)-1 {
				//Print each loop
				// 	fmt.Println(ln1)
				for j := 0; j < len(ln1); j++ {
					if j == len(ln1)-1 {
						fmt.Println(string(cMap[color]) + ln1[j] + cMap["reset"])
						continue
					}
					fmt.Print(string(cMap[color]) + ln1[j] + cMap["reset"])
				}

				for j := 0; j < len(ln2); j++ {
					if j == len(ln2)-1 {
						fmt.Println(string(cMap[color]) + ln2[j] + cMap["reset"])
						continue
					}
					fmt.Print(string(cMap[color]) + ln2[j] + cMap["reset"])
				}

				for j := 0; j < len(ln3); j++ {
					if j == len(ln3)-1 {
						fmt.Println(string(cMap[color]) + ln3[j] + cMap["reset"])
						continue
					}
					fmt.Print(string(cMap[color]) + ln3[j] + cMap["reset"])
				}

				for j := 0; j < len(ln4); j++ {
					if j == len(ln4)-1 {
						fmt.Println(string(cMap[color]) + ln4[j] + cMap["reset"])
						continue
					}
					fmt.Print(string(cMap[color]) + ln4[j] + cMap["reset"])
				}

				for j := 0; j < len(ln5); j++ {
					if j == len(ln5)-1 {
						fmt.Println(string(cMap[color]) + ln5[j] + cMap["reset"])
						continue
					}
					fmt.Print(string(cMap[color]) + ln5[j] + cMap["reset"])
				}

				for j := 0; j < len(ln6); j++ {
					if j == len(ln6)-1 {
						fmt.Println(string(cMap[color]) + ln6[j] + cMap["reset"])
						continue
					}
					fmt.Print(string(cMap[color]) + ln6[j] + cMap["reset"])
				}

				for j := 0; j < len(ln7); j++ {
					if j == len(ln7)-1 {
						fmt.Println(string(cMap[color]) + ln7[j] + cMap["reset"])
						continue
					}
					fmt.Print(string(cMap[color]) + ln7[j] + cMap["reset"])
				}

				for j := 0; j < len(ln8); j++ {
					if j == len(ln8)-1 {
						fmt.Println(string(cMap[color]) + ln8[j] + cMap["reset"])
						continue
					}
					fmt.Print(string(cMap[color]) + ln8[j] + cMap["reset"])
				}
			}
		}
	} else if len(os.Args) == 4 {
		num, errx := strconv.Atoi(os.Args[3])
		if errx != nil {
			log.Fatal(errx.Error())
		}

		for i := 0; i < len(SlcOutput); i++ {

			if i == dash && dash >= 0 {
				if len(ln1) != 0 && len(ln2) != 0 && len(ln3) != 0 && len(ln4) != 0 && len(ln5) != 0 && len(ln6) != 0 && len(ln7) != 0 && len(ln8) != 0 {

					for j := 0; j < len(ln1); j++ {
						if (j == num) && j == len(ln1)-1 {
							fmt.Println(string(cMap[color] + ln1[j] + string(cMap["reset"])))
							continue
						}

						if (j == num) {
							fmt.Print(string(cMap[color] + ln1[j] + string(cMap["reset"])))
							continue
						}

						if j == len(ln1)-1 {
							fmt.Println(ln1[j])
							continue
						}
						fmt.Print(ln1[j])
					}

					for j := 0; j < len(ln2); j++ {
						if (j == num) && j == len(ln2)-1 {
							fmt.Println(string(cMap[color] + ln2[j] + string(cMap["reset"])))
							continue
						}

						if (j == num) {
							fmt.Print(string(cMap[color] + ln2[j] + string(cMap["reset"])))
							continue
						}

						if j == len(ln2)-1 {
							fmt.Println(ln1[j])
							continue
						}
						fmt.Print(ln2[j])
					}

					for j := 0; j < len(ln3); j++ {
						if (j == num) && j == len(ln1)-1 {
							fmt.Println(string(cMap[color] + ln3[j] + string(cMap["reset"])))
							continue
						}

						if (j == num) {
							fmt.Print(string(cMap[color] + ln3[j] + string(cMap["reset"])))
							continue
						}

						if j == len(ln3)-1 {
							fmt.Println(ln3[j])
							continue
						}
						fmt.Print(ln3[j])
					}

					for j := 0; j < len(ln4); j++ {
						if (j == num) && j == len(ln4)-1 {
							fmt.Println(string(cMap[color] + ln4[j] + string(cMap["reset"])))
							continue
						}

						if (j== num) {
							fmt.Print(string(cMap[color] + ln4[j] + string(cMap["reset"])))
							continue
						}

						if j == len(ln4)-1 {
							fmt.Println(ln4[j])
							continue
						}
						fmt.Print(ln4[j])
					}

					for j := 0; j < len(ln5); j++ {
						if (j == num) && j == len(ln5)-1 {
							fmt.Println(string(cMap[color] + ln5[j] + string(cMap["reset"])))
							continue
						}

						if (j == num) {
							fmt.Print(string(cMap[color] + ln5[j] + string(cMap["reset"])))
							continue
						}

						if j == len(ln5)-1 {
							fmt.Println(ln5[j])
							continue
						}
						fmt.Print(ln5[j])
					}

					for j := 0; j < len(ln6); j++ {
						if (j == num) && j == len(ln6)-1 {
							fmt.Println(string(cMap[color] + ln6[j] + string(cMap["reset"])))
							continue
						}

						if (j == num) {
							fmt.Print(string(cMap[color] + ln6[j] + string(cMap["reset"])))
							continue
						}

						if j == len(ln6)-1 {
							fmt.Println(ln6[j])
							continue
						}
						fmt.Print(ln6[j])
					}

					for j := 0; j < len(ln7); j++ {
						if (j == num) && j == len(ln7)-1 {
							fmt.Println(string(cMap[color] + ln7[j] + string(cMap["reset"])))
							continue
						}

						if (j == num) {
							fmt.Print(string(cMap[color] + ln7[j] + string(cMap["reset"])))
							continue
						}

						if j == len(ln7)-1 {
							fmt.Println(ln7[j])
							continue
						}
						fmt.Print(ln7[j])
					}

					for j := 0; j < len(ln8); j++ {
						if (j == num) && j == len(ln8)-1 {
							fmt.Println(string(cMap[color] + ln8[j] + string(cMap["reset"])))
							continue
						}

						if (j == num) {
							fmt.Print(string(cMap[color] + ln8[j] + string(cMap["reset"])))
							continue
						}

						if j == len(ln8)-1 {
							fmt.Println(ln8[j])
							continue
						}
						fmt.Print(ln8[j])
					}

				} else {
					fmt.Println()
				}

				ln1 = []string{}
				ln2 = []string{}
				ln3 = []string{}
				ln4 = []string{}
				ln5 = []string{}
				ln6 = []string{}
				ln7 = []string{}
				ln8 = []string{}
				dash = dash + strings.Index(Output[dash+1:], `\n`) + 1
				i++
				num = num - i - 2
			} else {
				for j, s := range lttrlines {
					line = j
					if len(s) > 0 && s == string(SlcOutput[i]) {
						break
					}

				}

				for k := line + 1; k <= line+8; k++ {
					if k == line+1 {
						ln1 = append(ln1, lttrlines[k])
					}
					if k == line+2 {
						ln2 = append(ln2, lttrlines[k])
					}
					if k == line+3 {
						ln3 = append(ln3, lttrlines[k])
					}
					if k == line+4 {
						ln4 = append(ln4, lttrlines[k])
					}
					if k == line+5 {
						ln5 = append(ln5, lttrlines[k])
					}
					if k == line+6 {
						ln6 = append(ln6, lttrlines[k])
					}
					if k == line+7 {
						ln7 = append(ln7, lttrlines[k])
					}
					if k == line+8 {
						ln8 = append(ln8, lttrlines[k])
					}

				}
				if i == len(Output)-1 {
				
					//Print each loop
					// 	fmt.Println(ln1)
					for j := 0; j < len(ln1); j++ {
						if (j== num) && j == len(ln1)-1 {
							fmt.Println(string(cMap[color] + ln1[j] + string(cMap["reset"])))
							continue
						}

						if (j== num) {
							fmt.Print(string(cMap[color] + ln1[j] + string(cMap["reset"])))
							continue
						}

						if j == len(ln1)-1 {
							fmt.Println(ln1[j])
							continue
						}
						fmt.Print(ln1[j])
					}

					for j := 0; j < len(ln2); j++ {
						if (j== num) && j == len(ln2)-1 {
							fmt.Println(string(cMap[color] + ln2[j] + string(cMap["reset"])))
							continue
						}

						if (j== num) {
							fmt.Print(string(cMap[color] + ln2[j] + string(cMap["reset"])))
							continue
						}

						if j == len(ln2)-1 {
							fmt.Println(ln2[j])
							continue
						}
						fmt.Print(ln2[j])
					}

					for j := 0; j < len(ln3); j++ {
						if (j== num) && j == len(ln1)-1 {
							fmt.Println(string(cMap[color] + ln3[j] + string(cMap["reset"])))
							continue
						}

						if (j== num) {
							fmt.Print(string(cMap[color] + ln3[j] + string(cMap["reset"])))
							continue
						}

						if j == len(ln3)-1 {
							fmt.Println(ln3[j])
							continue
						}
						fmt.Print(ln3[j])
					}

					for j := 0; j < len(ln4); j++ {
						if (j== num) && j == len(ln4)-1 {
							fmt.Println(string(cMap[color] + ln4[j] + string(cMap["reset"])))
							continue
						}

						if (j== num) {
							fmt.Print(string(cMap[color] + ln4[j] + string(cMap["reset"])))
							continue
						}

						if j == len(ln4)-1 {
							fmt.Println(ln4[j])
							continue
						}
						fmt.Print(ln4[j])
					}

					for j := 0; j < len(ln5); j++ {
						if (j== num) && j == len(ln5)-1 {
							fmt.Println(string(cMap[color] + ln5[j] + string(cMap["reset"])))
							continue
						}

						if (j== num) {
							fmt.Print(string(cMap[color] + ln5[j] + string(cMap["reset"])))
							continue
						}

						if j == len(ln5)-1 {
							fmt.Println(ln5[j])
							continue
						}
						fmt.Print(ln5[j])
					}

					for j := 0; j < len(ln6); j++ {
						if (j== num) && j == len(ln6)-1 {
							fmt.Println(string(cMap[color] + ln6[j] + string(cMap["reset"])))
							continue
						}

						if (j== num) {
							fmt.Print(string(cMap[color] + ln6[j] + string(cMap["reset"])))
							continue
						}

						if j == len(ln6)-1 {
							fmt.Println(ln6[j])
							continue
						}
						fmt.Print(ln6[j])
					}

					for j := 0; j < len(ln7); j++ {
						if (j== num) && j == len(ln7)-1 {
							fmt.Println(string(cMap[color] + ln7[j] + string(cMap["reset"])))
							continue
						}

						if (j>= num) {
							fmt.Print(string(cMap[color] + ln7[j] + string(cMap["reset"])))
							continue
						}

						if j == len(ln7)-1 {
							fmt.Println(ln7[j])
							continue
						}
						fmt.Print(ln7[j])
					}

					for j := 0; j < len(ln8); j++ {
						if (j== num) && j == len(ln8)-1 {
							fmt.Println(string(cMap[color] + ln8[j] + string(cMap["reset"])))
							continue
						}

						if (j== num) {
							fmt.Print(string(cMap[color] + ln8[j] + string(cMap["reset"])))
							continue
						}

						if j == len(ln8)-1 {
							fmt.Println(ln8[j])
							continue
						}
						fmt.Print(ln8[j])
					}

				}
			}
				
			}
		} else if len(os.Args) == 5 {
			num1, err1 := strconv.Atoi(os.Args[3])
			if err1 != nil {
				log.Fatal(err1.Error())
			}
			num2, err2 := strconv.Atoi(os.Args[4])
			if err2 != nil {
				log.Fatal(err2.Error())
			}

		for i := 0; i < len(SlcOutput); i++ {

			if i == dash && dash >= 0 {
				if len(ln1) != 0 && len(ln2) != 0 && len(ln3) != 0 && len(ln4) != 0 && len(ln5) != 0 && len(ln6) != 0 && len(ln7) != 0 && len(ln8) != 0 {

					for j := 0; j < len(ln1); j++ {
						if (j>= num1 && j <= num2) && j == len(ln1)-1 {
							fmt.Println(string(cMap[color] + ln1[j] + string(cMap["reset"])))
							continue
						}

						if (j>= num1 && j <= num2) {
							fmt.Print(string(cMap[color] + ln1[j] + string(cMap["reset"])))
							continue
						}

						if j == len(ln1)-1 {
							fmt.Println(ln1[j])
							continue
						}
						fmt.Print(ln1[j])
					}

					for j := 0; j < len(ln2); j++ {
						if (j>= num1 && j <= num2) && j == len(ln2)-1 {
							fmt.Println(string(cMap[color] + ln2[j] + string(cMap["reset"])))
							continue
						}

						if (j>= num1 && j <= num2) {
							fmt.Print(string(cMap[color] + ln2[j] + string(cMap["reset"])))
							continue
						}

						if j == len(ln2)-1 {
							fmt.Println(ln1[j])
							continue
						}
						fmt.Print(ln2[j])
					}

					for j := 0; j < len(ln3); j++ {
						if (j>= num1 && j <= num2) && j == len(ln1)-1 {
							fmt.Println(string(cMap[color] + ln3[j] + string(cMap["reset"])))
							continue
						}

						if (j>= num1 && j <= num2) {
							fmt.Print(string(cMap[color] + ln3[j] + string(cMap["reset"])))
							continue
						}

						if j == len(ln3)-1 {
							fmt.Println(ln3[j])
							continue
						}
						fmt.Print(ln3[j])
					}

					for j := 0; j < len(ln4); j++ {
						if (j>= num1 && j <= num2) && j == len(ln4)-1 {
							fmt.Println(string(cMap[color] + ln4[j] + string(cMap["reset"])))
							continue
						}

						if (j>= num1 && j <= num2) {
							fmt.Print(string(cMap[color] + ln4[j] + string(cMap["reset"])))
							continue
						}

						if j == len(ln4)-1 {
							fmt.Println(ln4[j])
							continue
						}
						fmt.Print(ln4[j])
					}

					for j := 0; j < len(ln5); j++ {
						if (j>= num1 && j <= num2) && j == len(ln5)-1 {
							fmt.Println(string(cMap[color] + ln5[j] + string(cMap["reset"])))
							continue
						}

						if (j>= num1 && j <= num2) {
							fmt.Print(string(cMap[color] + ln5[j] + string(cMap["reset"])))
							continue
						}

						if j == len(ln5)-1 {
							fmt.Println(ln5[j])
							continue
						}
						fmt.Print(ln5[j])
					}

					for j := 0; j < len(ln6); j++ {
						if (j>= num1 && j <= num2) && j == len(ln6)-1 {
							fmt.Println(string(cMap[color] + ln6[j] + string(cMap["reset"])))
							continue
						}

						if (j>= num1 && j <= num2) {
							fmt.Print(string(cMap[color] + ln6[j] + string(cMap["reset"])))
							continue
						}

						if j == len(ln6)-1 {
							fmt.Println(ln6[j])
							continue
						}
						fmt.Print(ln6[j])
					}

					for j := 0; j < len(ln7); j++ {
						if (j>= num1 && j <= num2) && j == len(ln7)-1 {
							fmt.Println(string(cMap[color] + ln7[j] + string(cMap["reset"])))
							continue
						}

						if (j>= num1 && j <= num2) {
							fmt.Print(string(cMap[color] + ln7[j] + string(cMap["reset"])))
							continue
						}

						if j == len(ln7)-1 {
							fmt.Println(ln7[j])
							continue
						}
						fmt.Print(ln7[j])
					}

					for j := 0; j < len(ln8); j++ {
						if (j>= num1 && j <= num2) && j == len(ln8)-1 {
							fmt.Println(string(cMap[color] + ln8[j] + string(cMap["reset"])))
							continue
						}

						if (j>= num1 && j <= num2) {
							fmt.Print(string(cMap[color] + ln8[j] + string(cMap["reset"])))
							continue
						}

						if j == len(ln8)-1 {
							fmt.Println(ln8[j])
							continue
						}
						fmt.Print(ln8[j])
					}

				} else {
					fmt.Println()
				}

				ln1 = []string{}
				ln2 = []string{}
				ln3 = []string{}
				ln4 = []string{}
				ln5 = []string{}
				ln6 = []string{}
				ln7 = []string{}
				ln8 = []string{}
				dash = dash + strings.Index(Output[dash+1:], `\n`) + 1
				i++
				num1 = num1 - i - 2
				num2 = num2 - i - 2
			} else {
				for j, s := range lttrlines {
					line = j
					if len(s) > 0 && s == string(SlcOutput[i]) {
						break
					}

				}

				for k := line + 1; k <= line+8; k++ {
					if k == line+1 {
						ln1 = append(ln1, lttrlines[k])
					}
					if k == line+2 {
						ln2 = append(ln2, lttrlines[k])
					}
					if k == line+3 {
						ln3 = append(ln3, lttrlines[k])
					}
					if k == line+4 {
						ln4 = append(ln4, lttrlines[k])
					}
					if k == line+5 {
						ln5 = append(ln5, lttrlines[k])
					}
					if k == line+6 {
						ln6 = append(ln6, lttrlines[k])
					}
					if k == line+7 {
						ln7 = append(ln7, lttrlines[k])
					}
					if k == line+8 {
						ln8 = append(ln8, lttrlines[k])
					}

				}
				if i == len(Output)-1 {
				
					//Print each loop
					// 	fmt.Println(ln1)
					for j := 0; j < len(ln1); j++ {
						if (j>= num1 && j <= num2) && j == len(ln1)-1 {
							fmt.Println(string(cMap[color] + ln1[j] + string(cMap["reset"])))
							continue
						}

						if (j>= num1 && j <= num2) {
							fmt.Print(string(cMap[color] + ln1[j] + string(cMap["reset"])))
							continue
						}

						if j == len(ln1)-1 {
							fmt.Println(ln1[j])
							continue
						}
						fmt.Print(ln1[j])
					}

					for j := 0; j < len(ln2); j++ {
						if (j>= num1 && j <= num2) && j == len(ln2)-1 {
							fmt.Println(string(cMap[color] + ln2[j] + string(cMap["reset"])))
							continue
						}

						if (j>= num1 && j <= num2) {
							fmt.Print(string(cMap[color] + ln2[j] + string(cMap["reset"])))
							continue
						}

						if j == len(ln2)-1 {
							fmt.Println(ln2[j])
							continue
						}
						fmt.Print(ln2[j])
					}

					for j := 0; j < len(ln3); j++ {
						if (j>= num1 && j <= num2) && j == len(ln1)-1 {
							fmt.Println(string(cMap[color] + ln3[j] + string(cMap["reset"])))
							continue
						}

						if (j>= num1 && j <= num2) {
							fmt.Print(string(cMap[color] + ln3[j] + string(cMap["reset"])))
							continue
						}

						if j == len(ln3)-1 {
							fmt.Println(ln3[j])
							continue
						}
						fmt.Print(ln3[j])
					}

					for j := 0; j < len(ln4); j++ {
						if (j>= num1 && j <= num2) && j == len(ln4)-1 {
							fmt.Println(string(cMap[color] + ln4[j] + string(cMap["reset"])))
							continue
						}

						if (j>= num1 && j <= num2) {
							fmt.Print(string(cMap[color] + ln4[j] + string(cMap["reset"])))
							continue
						}

						if j == len(ln4)-1 {
							fmt.Println(ln4[j])
							continue
						}
						fmt.Print(ln4[j])
					}

					for j := 0; j < len(ln5); j++ {
						if (j>= num1 && j <= num2) && j == len(ln5)-1 {
							fmt.Println(string(cMap[color] + ln5[j] + string(cMap["reset"])))
							continue
						}

						if (j>= num1 && j <= num2) {
							fmt.Print(string(cMap[color] + ln5[j] + string(cMap["reset"])))
							continue
						}

						if j == len(ln5)-1 {
							fmt.Println(ln5[j])
							continue
						}
						fmt.Print(ln5[j])
					}

					for j := 0; j < len(ln6); j++ {
						if (j>= num1 && j <= num2) && j == len(ln6)-1 {
							fmt.Println(string(cMap[color] + ln6[j] + string(cMap["reset"])))
							continue
						}

						if (j>= num1 && j <= num2) {
							fmt.Print(string(cMap[color] + ln6[j] + string(cMap["reset"])))
							continue
						}

						if j == len(ln6)-1 {
							fmt.Println(ln6[j])
							continue
						}
						fmt.Print(ln6[j])
					}

					for j := 0; j < len(ln7); j++ {
						if (j>= num1 && j <= num2) && j == len(ln7)-1 {
							fmt.Println(string(cMap[color] + ln7[j] + string(cMap["reset"])))
							continue
						}

						if (j>= num1 && j <= num2) {
							fmt.Print(string(cMap[color] + ln7[j] + string(cMap["reset"])))
							continue
						}

						if j == len(ln7)-1 {
							fmt.Println(ln7[j])
							continue
						}
						fmt.Print(ln7[j])
					}

					for j := 0; j < len(ln8); j++ {
						if (j>= num1 && j <= num2) && j == len(ln8)-1 {
							fmt.Println(string(cMap[color] + ln8[j] + string(cMap["reset"])))
							continue
						}

						if (j>= num1 && j <= num2) {
							fmt.Print(string(cMap[color] + ln8[j] + string(cMap["reset"])))
							continue
						}

						if j == len(ln8)-1 {
							fmt.Println(ln8[j])
							continue
						}
						fmt.Print(ln8[j])
					}

				}
			}
				
			}


		}

		if err := scanner.Err(); err != nil {
					fmt.Println(err)
					// Handle the error
		}
		file.Close()
		fmt.Println(ln1)
		fmt.Println(ln2)
		fmt.Println(ln3)
		fmt.Println(ln4)
}


