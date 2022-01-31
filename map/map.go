package ascii

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func AsciiMap() {
	file, _ := ioutil.ReadFile("standard.txt")
	splitFile := strings.Split(string(file), "\n")
	fmt.Println(splitFile)
}
