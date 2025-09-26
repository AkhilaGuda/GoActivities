package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter file path")
	var filePath string
	fmt.Scanln(&filePath)
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error while opening file - %v\n", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	fmt.Println("File content: ")
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file %v\n", err)
	}

}
