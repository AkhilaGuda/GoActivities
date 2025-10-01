package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic", r)
		}
	}()
	nums := flag.String("nums", "", "comma-seperated numbers")
	flag.Parse()
	if *nums == "" {
		fmt.Println("Please enter numbers using -nums flag")
		os.Exit(1)
	}
	parts := strings.Split(*nums, ",")
	sum := 0
	for _, p := range parts {
		n, err := strconv.Atoi(p)
		if err != nil {
			panic(fmt.Sprintf("Invalid number: %s", p))

		}
		sum += n
	}
	fmt.Println("sum : ", sum)

}
