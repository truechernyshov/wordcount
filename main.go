package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	str := os.Args[1:]
	if str[0] == "" {
		fmt.Println(0)
	} else {
		a := strings.Split(str[0], " ")
		fmt.Println(len(a))
	}
}
