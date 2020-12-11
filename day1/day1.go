package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkIfSumExists(goal, numberToSkip int64) int64 {
	numberSkipped := false

	file, err := os.Open("day1/day1_input")
	check(err)

	// close the file after the function ends
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// numbers is a set
	numbers := map[int64]struct{}{}

	for scanner.Scan() {

		number, err := strconv.Atoi(scanner.Text())
		check(err)

		firstNumber := int64(number)

		if firstNumber != numberToSkip || numberSkipped {

			secondNumber := goal - firstNumber

			if _, containsKey := numbers[secondNumber]; containsKey {
				return firstNumber * secondNumber
			} else {
				numbers[firstNumber] = struct{}{}
			}
		} else {
			numberSkipped = true
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return -1
}

func main() {
	file, err := os.Open("day1/day1_input")
	check(err)

	defer file.Close()

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		number, err := strconv.Atoi(sc.Text())
		check(err)

		firstNumber := int64(number)
		goal := int64(2020 - number)

		result := checkIfSumExists(goal, firstNumber)

		if result != -1 {
			fmt.Println(result * firstNumber)
			os.Exit(0)
		}
	}
	fmt.Println("The are no such 3 numbers")
	os.Exit(-1)

}
