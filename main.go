package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	str := os.Args[1:]
	a := strings.Split(str[0], " ")
	fmt.Println(len(a))
}
