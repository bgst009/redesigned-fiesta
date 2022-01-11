package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	echo()
}

func echo() {
	var s, sep string = "", "-*-"
	for i := 1; i < len(os.Args); i++ {
		s += strconv.Itoa(i) + sep + os.Args[i] + "\n"
	}
	fmt.Println(s)
}
