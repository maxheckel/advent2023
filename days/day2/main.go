package main

import (
	"fmt"
	"github.com/maxheckel/advent2023/utils"
	"strings"
)

type Bag map[string]int

func (bag Bag) IsColorValid(color string, count int) bool {
	return count <= bag[color]
}

func (bag Bag) SetColorIfMax(color string, count int) {
	if bag[color] < count {
		bag[color] = count
	}
}

func (bag Bag) Power() int {
	return bag["red"] * bag["green"] * bag["blue"]
}

func main() {
	part1()
	part2()
}

func part2() {
	lines, err := utils.ReadInputLines("days/day2/input")
	if err != nil {
		fmt.Println(err)
	}
	sum := 0
	for _, line := range lines {
		sets := strings.Split(strings.Split(line, ": ")[1], ";")
		bag := Bag{}
		for _, set := range sets {
			for _, colorCount := range strings.Split(set, ",") {
				var color string
				var count int
				fmt.Sscanf(colorCount, "%d %s", &count, &color)
				bag.SetColorIfMax(color, count)
			}
		}
		sum += bag.Power()
	}
	fmt.Println(sum)
}

func part1() {
	bag := Bag{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	lines, err := utils.ReadInputLines("days/day2/input")
	if err != nil {
		fmt.Println(err)
	}

	sum := 0
TOP:
	for game, line := range lines {
		sets := strings.Split(strings.Split(line, ": ")[1], ";")
		for _, set := range sets {
			for _, colorCount := range strings.Split(set, ",") {
				var color string
				var count int
				fmt.Sscanf(colorCount, "%d %s", &count, &color)
				if !bag.IsColorValid(color, count) {
					continue TOP
				}
			}
		}
		sum += game + 1

	}
	fmt.Println(sum)
}
