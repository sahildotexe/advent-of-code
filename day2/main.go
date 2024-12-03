package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
	"math"
)

func main() {
	fi, err := os.Open("day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fi.Close()

	oi, err := os.Create("day2/output.txt")
	if err != nil {
		log.Fatalf("failed to create output: %v", err)
	}
	defer oi.Close()
	

	scanner := bufio.NewScanner(fi)
	safe_count := 0
	for scanner.Scan() {
		line := scanner.Text()
		report := strings.Fields(line)
		reportIsSafe := isSafeDampener(report)

		if reportIsSafe {
			_, err := oi.WriteString("Safe\n")
			if err != nil {
				log.Fatalf("failed to write output: %v", err)
			}
		} else {
			_, err := oi.WriteString("Unsafe\n")
			if err != nil {
				log.Fatalf("failed to write output: %v", err)
			}
		}

		if reportIsSafe {
			safe_count++
		}

	}
	fmt.Println(safe_count)
}

func isSafe(report []string) bool {
	nums := make([]float64, 0, len(report))
	for _, num := range report {
		f, err := strconv.ParseFloat(num, 64)
		if err != nil {
			log.Fatalf("failed to parse number: %v", err)
		}
		nums = append(nums, f)
	}

	asc := true
	dec := true
	diff := true

	for i := 0; i < len(nums) - 1; i++ {
		if nums[i] > nums[i+1] {
			asc = false
		}
		
		if nums[i] < nums[i+1] {
			dec = false
		}
		
		if nums[i] == nums[i+1] {
			diff = false
			asc = false
		}

		d := math.Abs(nums[i] - nums[i+1])

		if(d < 1 || d > 3) {
			diff = false
		}

		if !asc && !dec && !diff {
			return false
		}
	}

	return (asc && diff) || (dec && diff)
}

func isSafeDampener(report []string) bool {
	for i := 0; i < len(report); i++ {
		new_report := make([]string, len(report))
		copy(new_report, report)
		new_report = append(new_report[:i], new_report[i+1:]...)
		if isSafe(new_report) {
			return true
		}
	}
	return false
}

