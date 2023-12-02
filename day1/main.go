package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var textToInt map[string]string = map[string]string{
	"one":   "o1e",
	"two":   "t2o",
	"three": "t3e",
	"four":  "f4r",
	"five":  "f5e",
	"six":   "s6x",
	"seven": "s7n",
	"eight": "e8t",
	"nine":  "n9e",
}

func firstDigit(input string) int {
	for _, v := range input {
		if digit, err := strconv.Atoi(string(v)); err == nil {
			return digit
		}
	}
	return -1
}

func lastDigit(input string) int {
	for i := len(input) - 1; i >= 0; i-- {
		if digit, err := strconv.Atoi(string(input[i])); err == nil {
			return digit
		}
	}
	return -1
}

func replaceStrings(input string) string {
	for i := range input {
		for j := i + 1; j <= len(input); j++ {
			if _, ok := textToInt[input[i:j]]; ok {
				input = input[:i] + textToInt[input[i:j]] + input[j:]
			}
		}
	}
	return input
}

func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	contents := bufio.NewScanner(file)
	sum := 0

	for contents.Scan() {
		text := replaceStrings(contents.Text())

		first := firstDigit(text)
		last := lastDigit(text)

		number := first*10 + last
		sum += number
	}

	fmt.Println(sum)
}
