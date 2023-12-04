package main

import (
	"fmt"
	"github.com/maxheckel/advent2023/utils"
	"math"
	"strings"
)

func main() {
	part1()
	part2()
}

func part2() {
	lines, err := utils.ReadInputLines("days/day4/input")
	if err != nil {
		panic(err)
	}

	stack := make([]int, len(lines))
	for cardNum, line := range lines {
		stack[cardNum]++
		guessAndWinning := strings.Split(strings.Split(line, ": ")[1], "|")
		guess := strings.Fields(guessAndWinning[0])
		winning := strings.Fields(guessAndWinning[1])
		for x := 0; x < len(HashGeneric(guess, winning)); x++ {
			stack[cardNum+x+1] += stack[cardNum]
		}
	}
	sum := 0
	for _, count := range stack {
		sum += count
	}
	fmt.Println(sum)
}

func part1() {
	lines, err := utils.ReadInputLines("days/day4/input")
	if err != nil {
		panic(err)
	}

	sum := float64(0)
	for _, line := range lines {
		guessAndWinning := strings.Split(strings.Split(line, ": ")[1], "|")
		guess := strings.Fields(guessAndWinning[0])
		winning := strings.Fields(guessAndWinning[1])
		intersectCount := len(HashGeneric(guess, winning))
		if intersectCount > 0 {
			sum += math.Pow(float64(2), float64(intersectCount-1))
		}
	}
	fmt.Println(sum)
}

func HashGeneric[T comparable](a []T, b []T) []T {
	set := make([]T, 0)
	hash := make(map[T]struct{})

	for _, v := range a {
		hash[v] = struct{}{}
	}

	for _, v := range b {
		if _, ok := hash[v]; ok {
			set = append(set, v)
		}
	}

	return set
}
