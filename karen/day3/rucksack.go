package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	priority := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	total := 0

	file, err := os.Open("..\\..\\aoc_03_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Line: %v\n", line)

		a := line[:len(line)/2]
		b := line[len(line)/2:]
		fmt.Printf("a: %v\n", a)
		fmt.Printf("b: %v\n", b)

		c := commonItem(a, b)
		fmt.Printf("Common item: %v\n", c)

		p := strings.Index(priority, c) + 1
		fmt.Printf("Priority: %v\n", p)

		total += p

		fmt.Println("-------------------------------")
	}

	fmt.Printf("Total: %v\n", total)
}

func commonItem(a, b string) string {

	var c string

	for i := 0; i < len(b); i++ {
		if strings.Contains(a, b[i:i+1]) {
			c = b[i : i+1]
			break
		}
	}

	return c
}
