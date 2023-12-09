package main

import "fmt"

var input = map[int]int{
	53: 250,
	91: 1330,
	67: 1081,
	68: 1025,
}

var inputp2 = map[int]uint64{
	53916768: 250133010812015,
}

func main() {
	// part1()
	part2()
}

func part2() {
	// 1ms -> +1m/s
	result := 1
	for t, dist := range inputp2 {
		// charging
		cur := 0
		for i := 1; i < t; i++ {
			// charge for i seconds
			// then move (t-i) seconds
			// with i m/s speed
			curDist := uint64(i) * uint64(t-i)
			if curDist > dist {
				cur++
			}
		}
		result *= cur
	}
	fmt.Println(result)
}
func part1() {
	// 1ms -> +1m/s
	result := 1
	for t, dist := range input {
		// charging
		cur := 0
		for i := 1; i < t; i++ {
			// charge for i seconds
			// then move (t-i) seconds
			// with i m/s speed
			curDist := i * (t - i)
			if curDist > dist {
				cur++
			}
		}
		result *= cur
	}
	fmt.Println(result)
}
