package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
	"unicode"
)

func loadPuzzle() ([]string, error) {
	file, err := os.Open("input/day1.txt")
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Printf("failed to close file: %v", err)
		}
	}(file)

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func findFirstAndLastDigit(line string) (rune, rune) {
	var first, last rune

	foundFirst := false

	for _, char := range line {
		if unicode.IsDigit(char) {
			if !foundFirst {
				first = char
				foundFirst = true
			}

			last = char
		}
	}

	return first, last
}

func TestPartOne(t *testing.T) {
	lines, err := loadPuzzle()
	if err != nil {
		t.Fatalf("Failed to load puzzle: %v", err)
	}

	sum := uint64(0)

	for _, line := range lines {
		first, last := findFirstAndLastDigit(line)

		number, err := strconv.ParseUint(string([]rune{first, last}), 10, 64)
		if err != nil {
			t.Fatalf("Failed to convert string to uint: %v", err)
		}

		sum += number
	}

	t.Logf("%d", sum)
}

var prefixes = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"zero",
	"1",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
	"0",
}

func normalizePrefix(prefix string, index int) string {
	offset := len(prefixes) / 2

	if index < offset {
		return prefixes[index+offset]
	}

	return prefix
}

func findFirstAndLastNumber(line string) (string, string) {
	var first, last string

	foundFirst := false

	for i := 0; i < len(line); i++ {
		slice := line[i:]

		for p, prefix := range prefixes {
			if strings.HasPrefix(slice, prefix) {
				if !foundFirst {
					first = normalizePrefix(prefix, p)
					foundFirst = true
				}

				last = normalizePrefix(prefix, p)
				break
			}
		}
	}

	return first, last
}

func TestPartTwo(t *testing.T) {
	lines, err := loadPuzzle()
	if err != nil {
		t.Fatalf("Failed to load puzzle: %v", err)
	}

	sum := uint64(0)

	for _, line := range lines {
		first, last := findFirstAndLastNumber(line)

		number, err := strconv.ParseUint(first+last, 10, 64)
		if err != nil {
			t.Fatalf("Failed to convert string to uint: %v", err)
		}

		sum += number
	}

	t.Logf("%d", sum)
}
