package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var ans int

	file, err := os.Open("aoc_03_imput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		a := line[:len(line)/2]
		b := line[len(line)/2:]

		commonItem := commonItem(a, b)

		if commonItem >= 65 && commonItem <= 90 {
			commonItem -= 38
		} else if commonItem >= 97 && commonItem <= 122 {
			commonItem -= 96
		}

		ans += int(commonItem)
	}
	fmt.Println((ans))
}

func commonItem(a, b string) rune {
	for _, x := range a {
		for _, y := range b {

			if x == y {
				return x
			}
		}
	}
	return 0
}
