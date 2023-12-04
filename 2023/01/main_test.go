package main

import (
	"bufio"
	"strings"
	"testing"
)

func fakeReadLines(input string) <-chan string {
	lines := make(chan string, 10)
	go func() {
		defer close(lines)
		scanner := bufio.NewScanner(strings.NewReader(input))
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			lines <- scanner.Text()
		}
	}()
	return lines
}

func TestCalibrationPart1(t *testing.T) {
	input := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

	c, err := getCalibrationPart1(fakeReadLines(input))
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if c != 142 {
		t.Fatalf("expected 142, got %d", c)
	}
}

func TestCalibrationPart2(t *testing.T) {
	input := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

	c, err := getCalibrationPart2(fakeReadLines(input))
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if c != 281 {
		t.Fatalf("expected 281, got %d", c)
	}
}
