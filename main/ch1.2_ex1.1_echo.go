package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for i, v := range os.Args[0:] {
		fmt.Println(strconv.Itoa(i) + ":" + v)
	}
}
