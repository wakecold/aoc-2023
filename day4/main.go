package main

import (
	"bufio"
	"fmt"
	"os"
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
	m := map[int]int{} // card->mult
	index := 0
	for scanner.Scan() {
		line := scanner.Text()

		cards := strings.Split(line, ":")[1]
		cardInfo := strings.Split(cards, "|")
		winningCards := strings.Split(strings.Trim(cardInfo[0], " "), " ")
		myCards := strings.Split(strings.Trim(cardInfo[1], " "), " ")

		cur := 0
		for _, wCard := range winningCards {
			for _, mCard := range myCards {
				if wCard == "" {
					continue
				}
				if wCard == mCard {
					cur++
					break
				}
			}
		}
		m[index]++
		v, ok := m[index]
		if ok {
			for k := 0; k < v; k++ {
				for i := 1; i <= cur; i++ {
					m[index+i]++
				}
			}
		}
		index++
	}
	// Cards will never make you copy a card past the end of the table.
	result := 0
	for i := 0; i < index; i++ {
		if v, ok := m[i]; ok {
			result += v
		}
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

		cards := strings.Split(line, ":")[1]
		cardInfo := strings.Split(cards, "|")
		winningCards := strings.Split(strings.Trim(cardInfo[0], " "), " ")
		myCards := strings.Split(strings.Trim(cardInfo[1], " "), " ")

		cur := 0
		mult := 1
		// ugh n^2
		for _, wCard := range winningCards {
			for _, mCard := range myCards {
				// :/
				if wCard == "" {
					continue
				}
				if wCard == mCard {
					cur = mult
					mult *= 2
					break
				}
			}
		}
		result += cur
	}
	fmt.Println(result)
}
