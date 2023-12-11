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

func part1() {
	f, _ := os.Open("./input")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		vals := strings.Split(line, " ")
		arr := make([]int, len(vals))
		for i := range vals {
			val, err := strconv.Atoi(vals[i])
			if err != nil {
				panic(err)
			}
			arr[i] = val
		}
		result += predict(arr)
	}
	fmt.Println(result)
}

func part2() {
	f, _ := os.Open("./input")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		vals := strings.Split(line, " ")
		arr := make([]int, len(vals))
		for i := range vals {
			val, err := strconv.Atoi(vals[i])
			if err != nil {
				panic(err)
			}
			arr[i] = val
		}
		result += predictBackwards(arr)
	}
	fmt.Println(result)
}
func predict(arr []int) int {
	storage := make([][]int, 0)
	storage = append(storage, arr)
	for {
		cur := storage[len(storage)-1]
		if allZeroes(cur) {
			break
		}

		next := []int{}
		for i := 1; i < len(cur); i++ {
			next = append(next, cur[i]-cur[i-1])
		}
		storage = append(storage, next)
	}
	// go backwards and append each arr in storage
	// with prediction (take last val from prev arr in storage)
	// if we have the same value we just append with same value
	for i := len(storage) - 1; i >= 0; i-- {
		cur := storage[i]
		last := cur[len(cur)-1]
		prevLast := cur[len(cur)-2]
		if last == prevLast {
			cur = append(cur, last)
		} else {
			// take the value from prev storage
			prevCur := storage[i+1]
			cur = append(cur, last+prevCur[len(prevCur)-1])
		}
		storage[i] = cur
	}
	return storage[0][len(storage[0])-1]
}

func allZeroes(arr []int) bool {
	for i := range arr {
		if arr[i] != 0 {
			return false
		}
	}
	return true
}

func predictBackwards(arr []int) int {
	storage := make([][]int, 0)
	storage = append(storage, arr)
	for {
		cur := storage[len(storage)-1]
		if allZeroes(cur) {
			break
		}

		next := []int{}
		for i := len(cur) - 2; i >= 0; i-- {
			next = append([]int{cur[i+1] - cur[i]}, next...)
		}
		storage = append(storage, next)
	}
	for i := len(storage) - 1; i >= 0; i-- {
		cur := storage[i]
		first := cur[0]
		// if first == second {
		// ugh
		if i == len(storage)-1 || i == len(storage)-2 {
			cur = append([]int{first}, cur...)
		} else {
			prevCur := storage[i+1]
			cur = append([]int{first - prevCur[0]}, cur...)
		}
		storage[i] = cur
	}
	return storage[0][0]
}
