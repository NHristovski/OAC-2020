package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	//hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f

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

	fmt.Printf("The pairs %s HAVE ALL REQ KEYS\n", pairs)
	//requiredKeys := [7]string{"byr","iyr","eyr","hgt","hcl","ecl","pid"}
	if !validByr(pairs["byr"]) {
		fmt.Printf("WRONG BYR\n")
		return false
	}
	if !validIyr(pairs["iyr"]) {
		fmt.Printf("WRONG IYR\n")
		return false
	}
	if !validEyr(pairs["eyr"]) {
		fmt.Printf("WRONG EYR\n")
		return false
	}
	if !validHgt(pairs["hgt"]) {
		fmt.Printf("WRONG HGT\n")
		return false
	}
	if !validHcl(pairs["hcl"]) {
		fmt.Printf("WRONG HCL\n")
		return false
	}
	if !validEcl(pairs["ecl"]) {
		fmt.Printf("WRONG ECL\n")
		return false
	}
	if !validPid(pairs["pid"]) {
		fmt.Printf("WRONG PID\n")
		return false
	}

	fmt.Printf("The pairs %s are VALID\n", pairs)
	return true
}

func validPid(s string) bool {
	match, _ := regexp.MatchString("^[0-9]{9}$", s)
	return match
}

func validEcl(s string) bool {
	//	amb blu brn gry grn hzl oth
	match, _ := regexp.MatchString("^(amb|blu|brn|gry|grn|hzl|oth)$", s)
	return match
}

func validHcl(s string) bool {
	//hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f
	match, _ := regexp.MatchString("^#[0-9a-f]{6}$", s)
	return match
}

func validHgt(s string) bool {
	//hgt (Height) - a number followed by either cm or in:
	//
	//If cm, the number must be at least 150 and at most 193.
	//If in, the number must be at least 59 and at most 76.
	match, _ := regexp.MatchString("^(1[5-8][0-9]cm|19[0-3]cm|59in|6[0-9]in|7[0-6]in)$", s)
	return match
}

func validEyr(s string) bool {
	//eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	match, _ := regexp.MatchString("^(202[0-9]|2030)$", s)
	return match
}

func validIyr(s string) bool {
	//iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	match, _ := regexp.MatchString("^(201[0-9]|2020)$", s)
	return match
}

func validByr(s string) bool {
	// byr (Birth Year) - four digits; at least 1920 and at most 2002.
	match, _ := regexp.MatchString("^(19[2-9][0-9]|200[0-2])$", s)
	return match
}
