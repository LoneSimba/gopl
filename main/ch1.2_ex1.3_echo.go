package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	echo1(os.Args)
	fmt.Printf("%.2f elapsed\n", time.Since(start).Seconds())
	start = time.Now()
	echo2(os.Args)
	fmt.Printf("%.2f elapsed\n", time.Since(start).Seconds())
	start = time.Now()
	echo3(os.Args)
	fmt.Printf("%.2f elapsed\n", time.Since(start).Seconds())
}

func echo1(args []string) {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2(args []string) {
	s, sep := "", ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3(args []string) {
	fmt.Println(strings.Join(args[1:], " "))
}
