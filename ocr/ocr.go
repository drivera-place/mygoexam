package ocr

import (
	"errors"
	"fmt"
	"strings"
)

func checksum(line string) string {

	if strings.Contains(line, "?") {
		return "ILL"
	}

	multiplier := len(line)
	sum := 0
	for i := 1; i < len(line); i++ {
		sum += int(line[i]) * multiplier
		multiplier--
	}

	if sum%11 == 0 {
		return "OK"
	} else {
		return "ERR"
	}
}

func matchCase(digits map[int][]string) string {
	parsedString := ""

	zero := []string{" _ ", "| |", "|_|"}
	one := []string{"   ", "  |", "  |"}
	two := []string{" _ ", " _|", "|_ "}
	three := []string{" _ ", " _|", " _|"}
	four := []string{"   ", "|_|", "  |"}
	five := []string{" _ ", "|_ ", " _|"}
	six := []string{" _ ", "|_ ", "|_|"}
	seven := []string{" _ ", "  |", "  |"}
	eight := []string{" _ ", "|_|", "|_|"}
	nine := []string{" _ ", "|_|", " _|"}

	for i := 0; i < len(digits); i++ {

		if stringSlicesEqual(digits[i], zero) {
			parsedString += "0"
			continue
		}

		if stringSlicesEqual(digits[i], one) {
			parsedString += "1"
			continue
		}

		if stringSlicesEqual(digits[i], two) {
			parsedString += "2"
			continue
		}

		if stringSlicesEqual(digits[i], three) {
			parsedString += "3"
			continue
		}

		if stringSlicesEqual(digits[i], four) {
			parsedString += "4"
			continue
		}

		if stringSlicesEqual(digits[i], five) {
			parsedString += "5"
			continue
		}

		if stringSlicesEqual(digits[i], six) {
			parsedString += "6"
			continue
		}

		if stringSlicesEqual(digits[i], seven) {
			parsedString += "7"
			continue
		}

		if stringSlicesEqual(digits[i], eight) {
			parsedString += "8"
			continue
		}

		if stringSlicesEqual(digits[i], nine) {
			parsedString += "9"
			continue
		}

		parsedString += "?"
	}

	//fmt.Println(parsedString)

	result := checksum(parsedString)
	//fmt.Println(result)
	return parsedString + " " + result
}

func stringSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func ReadLine(s string) (string, error) {
	digits := map[int][]string{}

	lines := strings.Split(s, "\n")

	if !validate(lines) {
		return "", errors.New("Can not read digits.")
	}

	c := 0
	for i := 0; i < len(lines[0]); i = i + 3 {
		digits[c] = []string{lines[0][i : i+3], lines[1][i : i+3], lines[2][i : i+3]}
		c++
	}

	return matchCase(digits), nil
}

func validate(s []string) bool {
	return len(s) == 3 && ((len(s[0]) == 27) && (len(s[1]) == 27) && (len(s[2]) == 27))
}

func printDigital(n map[int][]string) {
	for _, v := range n {
		for _, i := range v {
			fmt.Println(i)
		}
	}
}
