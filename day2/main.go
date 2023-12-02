package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// part1()
	part2()
}

func part2() {
	f, _ := os.Open("./input")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		game := parts[1]
		sets := strings.Split(game, ";")
		cur := map[string]int{
			"red":   0,
			"blue":  0,
			"green": 0,
		}
		for _, set := range sets {
			cubes := strings.Split(set, ",")
			for _, cube := range cubes {
				// 3 blue
				vals := strings.Split(cube, " ")
				amount, err := strconv.Atoi(vals[1])
				if err != nil {
					panic(err)
				}
				color := vals[2]
				cur[color] = max(cur[color], amount)
			}
		}
		power := 1
		for _, v := range cur {
			power *= v
		}
		result += power
	}

	fmt.Println(result)
}
func part1() {
	// 12 red, 13 green, 14 blue
	rules := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	f, _ := os.Open("./input")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	result := 0
LineLoop:
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		game := parts[1]
		gameInfo := strings.Split(parts[0], " ")
		gameID, err := strconv.Atoi(gameInfo[len(gameInfo)-1])
		if err != nil {
			panic(err)
		}
		sets := strings.Split(game, ";")

		for _, set := range sets {
			cubes := strings.Split(set, ",")
			for _, cube := range cubes {
				// 3 blue
				vals := strings.Split(cube, " ")
				amount, err := strconv.Atoi(vals[1])
				if err != nil {
					panic(err)
				}
				color := vals[2]
				if ruleMax, ok := rules[color]; ok {
					if amount > ruleMax {
						continue LineLoop
					}
				} else {
					continue LineLoop
				}
			}
		}
		result += gameID
	}

	fmt.Println(result)
}
