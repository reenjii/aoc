package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const inputFile = "input.txt"

func main() {
	{
		lines, err := readLines(inputFile)
		if err != nil {
			log.Fatalf("fail to read input: %s", err)
		}

		part1, err := getCalibrationPart1(lines)
		if err != nil {
			log.Fatalf("fail to get calibration value part 1: %s", err)
		}
		fmt.Printf("calibration part 1 = %d\n", part1)
	}

	{
		lines, err := readLines(inputFile)
		if err != nil {
			log.Fatalf("fail to read input: %s", err)
		}

		part2, err := getCalibrationPart2(lines)
		if err != nil {
			log.Fatalf("fail to get calibration value part 2: %s", err)
		}
		fmt.Printf("calibration part 2 = %d\n", part2)
	}
}

func readLines(path string) (<-chan string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	lines := make(chan string, 10)

	// Read file in background
	go func() {
		defer file.Close()
		defer close(lines)

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			lines <- scanner.Text()
		}
	}()

	return lines, nil
}

func getCalibrationPart1(lines <-chan string) (int, error) {
	var sum int
	for line := range lines {
		calibration, err := getCalibrationValue(line)
		if err != nil {
			return 0, err
		}
		sum += calibration
	}
	return sum, nil
}

func getCalibrationPart2(lines <-chan string) (int, error) {
	var sum int
	for line := range lines {
		calibration, err := getCalibrationValueWithLetters(line)
		if err != nil {
			return 0, err
		}
		sum += calibration
	}
	return sum, nil
}

func getCalibrationValue(input string) (int, error) {
	var first, last int
	for i := 0; i < len(input); i++ {
		if first == 0 {
			if i := getDigit(input[i]); i > 0 {
				first = i
			}
		}
		if last == 0 {
			if i := getDigit(input[len(input)-i-1]); i > 0 {
				last = i
			}
		}
		if first != 0 && last != 0 {
			break
		}
	}
	if first == 0 || last == 0 {
		return 0, fmt.Errorf("fail to find first or last digit of input '%s'", input)
	}
	calibration, err := strconv.Atoi(strconv.Itoa(first) + strconv.Itoa(last))
	if err != nil {
		return 0, fmt.Errorf("fail to get calibration value of input '%s': %w", input, err)
	}
	return calibration, nil
}

var numbers = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func getCalibrationValueWithLetters(input string) (int, error) {
	var first, last int
	for i := 0; i < len(input); i++ {
		if first == 0 {
			if i := getDigitWithLetters(input[i:]); i > 0 {
				first = i
			}
		}
		if last == 0 {
			if i := getDigitWithLetters(input[len(input)-i-1:]); i > 0 {
				last = i
			}
		}
		if first != 0 && last != 0 {
			break
		}
	}
	if first == 0 || last == 0 {
		return 0, fmt.Errorf("fail to find first or last digit of input '%s'", input)
	}
	calibration, err := strconv.Atoi(strconv.Itoa(first) + strconv.Itoa(last))
	if err != nil {
		return 0, fmt.Errorf("fail to get calibration value of input '%s': %w", input, err)
	}
	return calibration, nil
}

func getDigit(input byte) int {
	i, err := strconv.Atoi(string(input))
	if err == nil {
		return i
	}
	return -1
}

func getDigitWithLetters(input string) int {
	// First look for a number written in letters
	for number, value := range numbers {
		l := len(number)
		if len(input) < l {
			// too short to fit number
			continue
		}
		if input[:l] == number {
			// found number in letters
			return value
		}
	}

	// Not found, fallback to getDigit
	return getDigit(input[0])
}
