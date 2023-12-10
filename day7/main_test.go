package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompare(t *testing.T) {
	a := "KK677"
	b := "32T3K"
	assert.True(t, compareHands(a, b))
}

func TestCompareSameRank(t *testing.T) {
	// a := "KK677"
	// b := "KTJJT"
	// assert.False(t, compareHands(a, b))
	a := "KTJJT"
	b := "KK677"
	assert.False(t, compareHands(a, b))
}

func TestHigh(t *testing.T) {
	a := "23678"
	b := "35274"
	assert.False(t, compareHands(a, b))
}

func TestSameRankWithJoker(t *testing.T) {
	a := "QQQJA"
	b := "KTJJT"
	assert.False(t, compareHandsWithJoker(a, b))
}

func TestFiveOfAKindWithJoker(t *testing.T) {
	a := "TTTJT"
	b := "TTTJJ"
	assert.True(t, compareHandsWithJoker(a, b))
}

func TestCorrectTypeWithJoker(t *testing.T) {
	a := "QQQJA"
	assert.True(t, isFourOfAKindWithJoker(a))
	b := "KTJJT"
	assert.True(t, isFourOfAKindWithJoker(b))

	a = "7JJJ2"
	assert.False(t, isFiveOfAKindWithJoker(a))
}
