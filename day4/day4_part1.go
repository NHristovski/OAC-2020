package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	const input_file = "day4/day4_input"

	file, err := os.Open(input_file)
	check(err)

	// close the file after the function ends
	defer file.Close()

	scanner := bufio.NewScanner(file)

	validPassportsCounter := 0

	pairs := map[string]string{}

	requiredKeys := [7]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			//fmt.Printf("#################### Empty line: %s\n",line)

			if validPassport(pairs, requiredKeys) {
				validPassportsCounter = validPassportsCounter + 1
				//fmt.Printf("The pairs %s are valid \n", pairs)
			} else {
				//fmt.Printf("The pairs %s are NOT valid \n", pairs)
			}
			pairs = map[string]string{}
		} else {
			//fmt.Printf("The line:\n%s\n is not empty\n",line)
			parts := strings.Split(line, " ")
			for _, part := range parts {
				subParts := strings.Split(part, ":")
				key := subParts[0]
				value := subParts[1]
				pairs[key] = value
			}
		}
	}

	if validPassport(pairs, requiredKeys) {
		validPassportsCounter = validPassportsCounter + 1
	}
	fmt.Println(validPassportsCounter)

}

func validPassport(pairs map[string]string, requiredKeys [7]string) bool {
	for _, key := range requiredKeys {
		_, ok := pairs[key]
		if !ok {
			fmt.Printf("The pairs %s are NOT valid because %s is missing\n", pairs, key)
			return false
		}
	}
	fmt.Printf("The pairs %s are VALID\n", pairs)
	return true
}
