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

type Node struct {
	left  *Node
	right *Node
	val   string
}

func part2() {
	f, _ := os.Open("./input")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	// line 1 - template path
	// line 2 - empty
	// line 3... - connections

	idx := 0
	instruction := ""
	nodeStorage := map[string]*Node{}
	for scanner.Scan() {
		line := scanner.Text()
		if idx == 0 {
			// there is a space in the instruction
			instruction = line
			idx++
			continue
		}
		if line == "" {
			continue
		}

		info := strings.Split(line, "=")
		cur := strings.Trim(info[0], " ")
		child := strings.Split(info[1], ",")
		left := strings.Trim(child[0], " (")
		right := strings.Trim(child[1], ")")
		right = strings.Trim(right, " ")

		var node *Node
		if n, ok := nodeStorage[cur]; ok {
			node = n
		} else {
			n := &Node{nil, nil, cur}
			nodeStorage[cur] = n
			node = n
		}
		if l, ok := nodeStorage[left]; ok {
			node.left = l
		} else {
			l := &Node{nil, nil, left}
			nodeStorage[left] = l
			node.left = l
		}
		if r, ok := nodeStorage[right]; ok {
			node.right = r
		} else {
			r := &Node{nil, nil, right}
			nodeStorage[right] = r
			node.right = r
		}
	}
	starts := []string{}
	for k := range nodeStorage {
		for _, c := range k {
			if c == 'A' {
				starts = append(starts, k)
				break
			}
		}
	}
	i := 0
	vals := []int{}
	for _, start := range starts {
		curVal := 0

		cur := nodeStorage[start]
	StartLoop:
		for {
			if i == len(instruction) {
				i = 0
			}

			if cur == nil {
				panic("cur node is nil")
			}

			for _, c := range cur.val {
				if c == 'Z' {
					break StartLoop
				}
			}

			switch instruction[i] {
			case 'L':
				cur = cur.left
			case 'R':
				cur = cur.right
			}
			i++
			curVal++
		}
		vals = append(vals, curVal)
	}
	fmt.Println(lcm(vals[0], vals[1], vals[2:]...))
}
func part1() {
	f, _ := os.Open("./input")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	// line 1 - template path
	// line 2 - empty
	// line 3... - connections

	idx := 0
	instruction := ""
	nodeStorage := map[string]*Node{}
	for scanner.Scan() {
		line := scanner.Text()
		if idx == 0 {
			// there is a space in the instruction
			instruction = line
			idx++
			continue
		}
		if line == "" {
			continue
		}

		info := strings.Split(line, "=")
		cur := strings.Trim(info[0], " ")
		child := strings.Split(info[1], ",")
		left := strings.Trim(child[0], " (")
		right := strings.Trim(child[1], ")")
		right = strings.Trim(right, " ")

		var node *Node
		if n, ok := nodeStorage[cur]; ok {
			node = n
		} else {
			n := &Node{nil, nil, cur}
			nodeStorage[cur] = n
			node = n
		}
		if l, ok := nodeStorage[left]; ok {
			node.left = l
		} else {
			l := &Node{nil, nil, left}
			nodeStorage[left] = l
			node.left = l
		}
		if r, ok := nodeStorage[right]; ok {
			node.right = r
		} else {
			r := &Node{nil, nil, right}
			nodeStorage[right] = r
			node.right = r
		}
	}
	i := 0
	start, finish := "AAA", "ZZZ"
	cur := nodeStorage[start]
	result := 0
	for {
		if i == len(instruction) {
			i = 0
		}

		if cur == nil {
			panic("cur node is nil")
		}

		if cur.val == finish {
			break
		}

		switch instruction[i] {
		case 'L':
			cur = cur.left
		case 'R':
			cur = cur.right
		}
		i++
		result++
	}
	fmt.Println(result)
}

// source - random google link
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)
	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}
	return result
}
