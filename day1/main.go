package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	part1()
	// part2()
}

func part2() {
	f, _ := os.Open("./input")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	result := 0

	digits := map[string]int{
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

	for scanner.Scan() {
		line := scanner.Text()
		cur := 0
		// left
	ForwardLineLoop:
		for i := range line {
			if isNumber(line[i]) {
				n, _ := strconv.Atoi(string(line[i]))
				cur = 10 * n
				break
			} else {
				for d, v := range digits {
					if i+len(d) <= len(line) && line[i:i+len(d)] == d {
						cur = 10 * v
						break ForwardLineLoop
					}
				}
			}
		}

		// right
	BackwardLineLoop:
		for i := len(line) - 1; i >= 0; i-- {
			if isNumber(line[i]) {
				n, _ := strconv.Atoi(string(line[i]))
				cur += n
				break
			} else {
				for d, v := range digits {
					if i-len(d) >= 0 && line[i-len(d)+1:i+1] == d {
						cur += v
						break BackwardLineLoop
					}
				}
			}
		}
		result += cur
	}
	fmt.Println(result)
}

func part1() {

	f, _ := os.Open("./input")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	result := 0

	for scanner.Scan() {
		line := scanner.Text()
		cur := 0
		// left
		for i := range line {
			if isNumber(line[i]) {
				n, _ := strconv.Atoi(string(line[i]))
				cur = 10 * n
				break
			}
		}

		// right
		for i := len(line) - 1; i >= 0; i-- {
			if isNumber(line[i]) {
				n, _ := strconv.Atoi(string(line[i]))
				cur += n
				break
			}
		}
		result += cur
	}
	fmt.Println(result)
}

func isNumber(c byte) bool {
	return int(c) >= int('0') && int(c) <= int('9')
}
