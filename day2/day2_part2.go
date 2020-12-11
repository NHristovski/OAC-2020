package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	const input_file = "day2/day2_input"

	file, err := os.Open(input_file)
	check(err)

	// close the file after the function ends
	defer file.Close()

	scanner := bufio.NewScanner(file)

	const separator = " "

	validPasswordsCounter := 0

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, separator)

		minRange, maxRange := getRanges(parts[0])

		letter := parts[1][0]

		password := parts[2]

		if (password[minRange-1] == letter && password[maxRange-1] != letter) ||
			(password[minRange-1] != letter && password[maxRange-1] == letter) {
			validPasswordsCounter = validPasswordsCounter + 1
		}
		//fmt.Printf("MinRange: %d , MaxRange: %d , Letter %s ,  password: %s \n", minRange, maxRange, letter, password)
	}
	fmt.Println(validPasswordsCounter)
}

func getRanges(ranges string) (int, int) {
	rangesParts := strings.Split(ranges, "-")

	minRange, err := strconv.Atoi(rangesParts[0])
	check(err)
	maxRange, err := strconv.Atoi(rangesParts[1])
	check(err)

	return minRange, maxRange
}
