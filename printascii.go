package ascii

import (
	"fmt"
	"strings"
)

func Createmap(slice []string) map[int][]string {
	mapslice := make(map[int][]string)
	ascii := 32
	count := 0
	split := 0
	empty := []string{}

	for split < len(slice)-9 {
		split = count * 9
		for i := split; i < split+9; i++ {
			empty = append(empty, slice[i])
		}
		mapslice[ascii] = empty
		empty = []string{}
		ascii++
		count++
	}

	return mapslice
}

func Printascii(mapascii map[int][]string, word string) string {
	toprint := ""
	for i := 1; i <= 7; i++ {
		for _, c := range word {
			toprint += mapascii[int(c)][i]
		}
		toprint += "\n"
	}
	return toprint
}

func Condition(mapascii map[int][]string, a string) {
	splitA := strings.Split(a, "\\n")
	for _, w := range splitA {
		if w != "" {
			fmt.Print(Printascii(mapascii, w))
		} else {
			fmt.Println()
		}
	}
}
