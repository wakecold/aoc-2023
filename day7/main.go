package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var handRanks = map[byte]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'J': 10,
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
}

var handRanksWithJoker = map[byte]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
	'J': 1,
}

type Hand struct {
	hand string
	bet  int
}

func main() {
	// part1()
	part2()
}

func part2() {
	f, _ := os.Open("./input")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	hands := []Hand{}
	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(line, " ")
		hand := data[0]
		bet, err := strconv.Atoi(data[1])
		if err != nil {
			panic(err)
		}
		hands = append(hands, Hand{hand, bet})
	}

	sort.Slice(hands, func(i, j int) bool {
		return compareHandsWithJoker(hands[i].hand, hands[j].hand)
	})
	var result uint64

	for i := len(hands) - 1; i >= 0; i-- {
		rank := len(hands) - i
		cur := hands[i]
		result += uint64(cur.bet) * uint64(rank)
	}
	fmt.Println(result)
}

// returns true if a>b
func compareHandsWithJoker(a, b string) bool {
	// five of a kind
	if isFiveOfAKindWithJoker(a) {
		if isFiveOfAKindWithJoker(b) {
			return orderCompareWithJoker(a, b)
		}
		return true
	}
	if isFiveOfAKindWithJoker(b) {
		return false
	}
	// four of a kind
	if isFourOfAKindWithJoker(a) {
		if isFourOfAKindWithJoker(b) {
			return orderCompareWithJoker(a, b)
		}
		return true
	}
	if isFourOfAKindWithJoker(b) {
		return false
	}
	// fullhouse, 3+2
	if isFullHouseWithJoker(a) {
		if isFullHouseWithJoker(b) {
			return orderCompareWithJoker(a, b)
		}
		return true
	}
	if isFullHouseWithJoker(b) {
		return false
	}

	// three of a kind
	if isThreeOfAKindWithJoker(a) {
		if isThreeOfAKindWithJoker(b) {
			return orderCompareWithJoker(a, b)
		}
		return true
	}
	if isThreeOfAKindWithJoker(b) {
		return false
	}

	// two pair
	if isTwoPairWithJoker(a) {
		if isTwoPairWithJoker(b) {
			return orderCompareWithJoker(a, b)
		}
		return true
	}
	if isTwoPairWithJoker(b) {
		return false
	}

	// one pair
	if isPairWithJoker(a) {
		if isPairWithJoker(b) {
			return orderCompareWithJoker(a, b)
		}
		return true
	}
	if isPairWithJoker(b) {
		return false
	}

	return orderCompareWithJoker(a, b)
}

func hasJoker(a string) bool {
	for i := range a {
		if a[i] == 'J' {
			return true
		}
	}
	return false
}

func countJokers(a string) int {
	result := 0
	for i := range a {
		if a[i] == 'J' {
			result++
		}
	}
	return result
}

func isFiveOfAKindWithJoker(a string) bool {
	if isFiveOfAKind(a) {
		return true
	}
	if !hasJoker(a) {
		return false
	}
	count := countJokers(a)
	if count == 1 {
		return isFourOfAKind(a)
	}
	if count == 2 {
		return isThreeOfAKind(a)
	}
	if count == 3 {
		//77JJJ
		b := ""
		for i := range a {
			if a[i] != 'J' {
				b += string(a[i])
			}
		}
		if len(b) != 2 {
			return false
		}
		return b[0] == b[1]
	}
	// count==4
	return true
}

func isFourOfAKindWithJoker(a string) bool {
	if isFourOfAKind(a) {
		return true
	}
	if !hasJoker(a) {
		return false
	}
	// JJJKQ
	// KKKJQ
	// ???
	count := countJokers(a)
	if count == 1 {
		return isThreeOfAKind(a)
	}
	if count == 2 {
		// JJKKQ
		b := ""
		for i := range a {
			if a[i] != 'J' {
				b += string(a[i])
			}
		}
		return isPair(b)
	}
	// count==3
	return true
}

func isThreeOfAKindWithJoker(a string) bool {
	if isThreeOfAKind(a) {
		return true
	}
	if !hasJoker(a) {
		return false
	}
	count := countJokers(a)
	if count == 1 {
		return isPair(a)
	}
	// count==2
	return true
}

func isFullHouseWithJoker(a string) bool {
	if isFullHouse(a) {
		return true
	}
	if !hasJoker(a) {
		return false
	}
	// JJKKQ
	// KKQQJ
	count := countJokers(a)
	if count == 1 {
		return isTwoPair(a)
	}
	if count == 2 {
		b := ""
		for i := range a {
			if a[i] != 'J' {
				b += string(a[i])
			}
		}
		return isPair(b)
	}
	// count == 3
	// JJJKQ -> KJJ QJ
	return true
}

func isTwoPairWithJoker(a string) bool {
	if isTwoPair(a) {
		return true
	}
	if !hasJoker(a) {
		return false
	}
	// JJKQA
	// KKQAJ
	count := countJokers(a)
	if count == 1 {
		return isPair(a)
	}
	// JJKQA -> KJ QJ A
	return true
}

func isPairWithJoker(a string) bool {
	if isPair(a) {
		return true
	}
	return hasJoker(a)
}

func orderCompareWithJoker(a, b string) bool {
	for i := range a {
		if a[i] == b[i] {
			continue
		}
		return handRanksWithJoker[a[i]] > handRanksWithJoker[b[i]]
	}
	panic("equal hands")
}

func part1() {
	f, _ := os.Open("./input")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	hands := []Hand{}
	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(line, " ")
		hand := data[0]
		bet, err := strconv.Atoi(data[1])
		if err != nil {
			panic(err)
		}
		hands = append(hands, Hand{hand, bet})
	}
	sort.Slice(hands, func(i, j int) bool {
		return compareHands(hands[i].hand, hands[j].hand)
	})
	var result uint64

	for i := len(hands) - 1; i >= 0; i-- {
		rank := len(hands) - i
		cur := hands[i]
		result += uint64(cur.bet) * uint64(rank)
	}
	fmt.Println(result)
}

//	A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2
//
// return true if a>b
func compareHands(a, b string) bool {
	// 32T3K vs T55J5
	// five of a kind
	if isFiveOfAKind(a) {
		if isFiveOfAKind(b) {
			return handRanks[a[0]] > handRanks[b[0]]
		}
		return true
	}
	if isFiveOfAKind(b) {
		return false
	}

	// four of a kind
	if isFourOfAKind(a) {
		if isFourOfAKind(b) {
			return orderCompare(a, b)
		}
		return true
	}
	if isFourOfAKind(b) {
		return false
	}

	// fullhouse, 3+2
	if isFullHouse(a) {
		if isFullHouse(b) {
			return orderCompare(a, b)
		}
		return true
	}
	if isFullHouse(b) {
		return false
	}

	// three of a kind
	if isThreeOfAKind(a) {
		if isThreeOfAKind(b) {
			return orderCompare(a, b)
		}
		return true
	}
	if isThreeOfAKind(b) {
		return false
	}

	// two pair
	if isTwoPair(a) {
		if isTwoPair(b) {
			return orderCompare(a, b)
		}
		return true
	}
	if isTwoPair(b) {
		return false
	}

	// one pair
	if isPair(a) {
		if isPair(b) {
			return orderCompare(a, b)
		}
		return true
	}
	if isPair(b) {
		return false
	}

	// high card
	return orderCompare(a, b)
}

// true if a > b
func orderCompare(a, b string) bool {
	for i := range a {
		if a[i] == b[i] {
			continue
		}
		return handRanks[a[i]] > handRanks[b[i]]
	}
	panic("equal hands")
}

func isFiveOfAKind(a string) bool {
	a0 := a[0]
	for i := range a {
		if a[i] != a0 {
			return false
		}
	}
	return true
}

func isFourOfAKind(a string) bool {
	freq := map[byte]int{}
	for i := range a {
		freq[a[i]]++
	}
	for _, v := range freq {
		if v == 4 {
			return true
		}
	}
	return false
}

func isThreeOfAKind(a string) bool {
	freq := map[byte]int{}
	for i := range a {
		freq[a[i]]++
	}

	for _, v := range freq {
		if v == 3 {
			return true
		}
	}
	return false
}

func isFullHouse(a string) bool {
	freq := map[byte]int{}
	for i := range a {
		freq[a[i]]++
	}
	if len(freq) != 2 {
		return false
	}
	for _, v := range freq {
		if v != 2 && v != 3 {
			return false
		}
	}
	return true
}

func isTwoPair(a string) bool {
	freq := map[byte]int{}
	for i := range a {
		freq[a[i]]++
	}

	// 2+2+1
	if len(freq) != 3 {
		return false
	}

	for _, v := range freq {
		if v != 2 && v != 1 {
			return false
		}
	}
	return true
}

func isPair(a string) bool {
	for i := range a {
		for k := i + 1; k < len(a); k++ {
			if a[k] == a[i] {
				return true
			}
		}
	}
	return false
}
