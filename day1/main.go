package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	firstNums := make([]int, 0, 1000)
	secondNums := make([]int, 0, 1000)

	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Fields(line)

		first, err := strconv.Atoi(nums[0])
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		second, err := strconv.Atoi(nums[1])
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		firstNums = append(firstNums, first)
		secondNums = append(secondNums, second)
	}

	sort.Ints(firstNums)
	sort.Ints(secondNums)

	totalDiff := 0
	for i := range firstNums {
		diff := firstNums[i] - secondNums[i]
		if diff < 0 {
			diff = -diff
		}
		totalDiff += diff
	}
	fmt.Println("Part 1:", totalDiff)

	freqs := make(map[int]int)
	for _, num := range secondNums {
		freqs[num]++
	}

	similarityScore := 0
	for _, num := range firstNums {
		similarityScore += num * freqs[num]
	}
	fmt.Println("Part 2:", similarityScore)
}
