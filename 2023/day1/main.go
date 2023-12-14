package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// "2023/day1/example.txt"
func main() {
	lines, err := GetLinesFromFile("2023/day1/inputp1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	res := SummarizeLines(lines, GetNumbersPartTwo)
	fmt.Println(res)
}

// Part 2
func GetNumbersPartTwo(line string) int {
	// one, two, three, four, five, six, seven, eight, and nine
	numberMatch := map[string]string{
		"one":   "one1one",
		"two":   "two2two",
		"three": "three3three",
		"four":  "four4four",
		"five":  "five5five",
		"six":   "six6six",
		"seven": "seven7seven",
		"eight": "eight8eight",
		"nine":  "nine9nine",
	}
	clstr := line
	for k, v := range numberMatch {
		val := strings.ReplaceAll(clstr, k, v)
		clstr = val
		fmt.Println(clstr)
	}
	res := GetNumbersPartOne(clstr)
	fmt.Println(res)
	return res
}

// Part 1
func SummarizeLines(lines []string, aggfunc func(line string) int) int {
	result := 0
	for _, l := range lines {
		num := aggfunc(l)
		result += num
	}
	return result
}

func GetNumbersPartOne(line string) int {
	var res int
	i := 0
	j := len(line) - 1
	for i <= j {
		first, err := strconv.Atoi(string(line[i]))
		if err != nil {
			i++
			continue
		}
		last, err := strconv.Atoi(string(line[j]))
		if err != nil {
			j--
			continue
		}
		snum := fmt.Sprintf("%v%v", first, last)
		num, _ := strconv.Atoi(snum)
		res = num
		break
	}
	return res
}

func GetLinesFromFile(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	lines := []string{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}
