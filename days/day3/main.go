package main

import (
	"fmt"
	"github.com/maxheckel/advent2023/utils"
	"strconv"
	"unicode"
)

func isSymbol(r rune) bool {
	if unicode.IsNumber(r) {
		return false
	}
	if r == '.' {
		return false
	}
	return true
}

func main() {
	part1()
	part2()
}

func part2() {
	grid := buildGrid()
	sum := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == '*' {
				sum += findGearRatio(y, x, grid)
			}
		}
	}
	fmt.Println(sum)
}

func part1() {
	grid := buildGrid()

	seenGrid := [][]bool{}
	for _, r := range grid {
		seenGrid = append(seenGrid, make([]bool, len(r)))
	}
	sum := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if isSymbol(grid[y][x]) {
				sum += findUnseenAdjacentNumbersSum(y, x, grid, seenGrid)
			}
		}

	}
	fmt.Println(sum)
}

func findGearRatio(y, x int, grid [][]rune) int {
	nums := []int{}
	maxX := len(grid[y]) - 1
	maxY := len(grid) - 1
	seenGrid := [][]bool{}
	for _, r := range grid {
		seenGrid = append(seenGrid, make([]bool, len(r)))
	}
	if x-1 >= 0 && y-1 >= 0 {
		if !seenGrid[y-1][x-1] {
			num := getFullNumber(y-1, x-1, grid, seenGrid)
			if num > 0 {
				nums = append(nums, num)
			}
		}
	}
	if x+1 <= maxX && y+1 <= maxY {
		if !seenGrid[y+1][x+1] {
			num := getFullNumber(y+1, x+1, grid, seenGrid)
			if num > 0 {
				nums = append(nums, num)
			}
		}
	}
	if x-1 >= 0 && y+1 <= maxY {
		if !seenGrid[y+1][x-1] {
			num := getFullNumber(y+1, x-1, grid, seenGrid)
			if num > 0 {
				nums = append(nums, num)
			}
		}
	}
	if x+1 <= maxX && y-1 >= 0 {
		if !seenGrid[y-1][x+1] {
			num := getFullNumber(y-1, x+1, grid, seenGrid)
			if num > 0 {
				nums = append(nums, num)
			}
		}
	}
	if x-1 >= 0 {
		if !seenGrid[y][x-1] {
			num := getFullNumber(y, x-1, grid, seenGrid)
			if num > 0 {
				nums = append(nums, num)
			}
		}
	}
	if x+1 <= maxX {
		if !seenGrid[y][x+1] {
			num := getFullNumber(y, x+1, grid, seenGrid)
			if num > 0 {
				nums = append(nums, num)
			}
		}
	}
	if y-1 >= 0 {
		if !seenGrid[y-1][x] {
			num := getFullNumber(y-1, x, grid, seenGrid)
			if num > 0 {
				nums = append(nums, num)
			}
		}
	}
	if y+1 <= maxY {
		if !seenGrid[y+1][x] {
			num := getFullNumber(y+1, x, grid, seenGrid)
			if num > 0 {
				nums = append(nums, num)
			}
		}
	}

	if len(nums) >= 2 {
		product := 1
		for _, num := range nums {
			product *= num
		}
		return product
	}
	return 0
}

func findUnseenAdjacentNumbersSum(y, x int, grid [][]rune, seenGrid [][]bool) int {
	sum := 0
	maxX := len(grid[y]) - 1
	maxY := len(grid) - 1

	if x-1 >= 0 && y-1 >= 0 {
		if !seenGrid[y-1][x-1] {
			sum += getFullNumber(y-1, x-1, grid, seenGrid)
		}
	}
	if x+1 <= maxX && y+1 <= maxY {
		if !seenGrid[y+1][x+1] {
			sum += getFullNumber(y+1, x+1, grid, seenGrid)
		}
	}
	if x-1 >= 0 && y+1 <= maxY {
		if !seenGrid[y+1][x-1] {
			sum += getFullNumber(y+1, x-1, grid, seenGrid)
		}
	}
	if x+1 <= maxX && y-1 >= 0 {
		if !seenGrid[y-1][x+1] {
			sum += getFullNumber(y-1, x+1, grid, seenGrid)
		}
	}
	if x-1 >= 0 {
		if !seenGrid[y][x-1] {
			sum += getFullNumber(y, x-1, grid, seenGrid)
		}
	}
	if x+1 <= maxX {
		if !seenGrid[y][x+1] {
			sum += getFullNumber(y, x+1, grid, seenGrid)
		}
	}
	if y-1 >= 0 {
		if !seenGrid[y-1][x] {
			sum += getFullNumber(y-1, x, grid, seenGrid)
		}
	}
	if y+1 <= maxY {
		if !seenGrid[y+1][x] {
			sum += getFullNumber(y+1, x, grid, seenGrid)
		}
	}
	return sum
}

func getFullNumber(y, x int, grid [][]rune, seenGrid [][]bool) int {
	next := x
	before := ""
	if !unicode.IsNumber(grid[y][x]) {
		return 0
	}
	for true {
		if next-1 >= 0 {
			next--
		} else {
			break
		}
		if !unicode.IsNumber(grid[y][next]) {
			break
		}
		seenGrid[y][next] = true
		before = string(grid[y][next]) + before
	}
	next = x
	after := ""
	for true {
		if next+1 <= len(grid[y])-1 {
			next++
		} else {
			break
		}
		if !unicode.IsNumber(grid[y][next]) {
			break
		}
		seenGrid[y][next] = true
		after += string(grid[y][next])
	}
	// We should always have at least one number here
	num, _ := strconv.Atoi(before + string(grid[y][x]) + after)
	return num
}

func buildGrid() [][]rune {
	lines, err := utils.ReadInputLines("days/day3/input")
	if err != nil {
		panic(err)
	}
	grid := [][]rune{}
	for y, line := range lines {
		grid = append(grid, []rune{})
		for _, char := range line {
			grid[y] = append(grid[y], char)
		}
	}
	return grid
}
