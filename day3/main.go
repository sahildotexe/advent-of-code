package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	b, err := os.ReadFile("day3/input.txt")
	if err != nil {
		fmt.Print(err)
	}

	oi, err := os.Create("day3/output.txt")
	if err != nil {
		log.Fatalf("failed to create output: %v", err)
	}
	defer oi.Close()

	s := string(b)
	res := evaluate(s)
	res2 := evaluateWithDo(s)
	oi.WriteString(fmt.Sprintf("%d\n%d\n", res, res2))
}

func evaluate(s string) int {
	n := len(s)
	res := 0

	for i := 0; i < n-3; i++ {
		if s[i] == 'm' && s[i+1] == 'u' && s[i+2] == 'l' && s[i+3] == '(' {
			num1 := ""
			isNum1 := true
			num2 := ""
			isNum2 := true
			for j := i + 4; j < n; j++ {
				if isDigit(s[j]) {
					num1 += string(s[j])
				} else if s[j] == ',' {
					break
				} else {
					isNum1 = false
					break
				}
			}

			if isNum1 {
				for j := i + 4 + len(num1) + 1; j < n; j++ {
					if isDigit(s[j]) {
						num2 += string(s[j])
					} else if s[j] == ')' {
						break
					} else {
						isNum2 = false
						break
					}
				}
			}

			if isNum1 && isNum2 {
				n1, err := strconv.Atoi(num1)
				if err != nil {
					continue
				}
				n2, err := strconv.Atoi(num2)
				if err != nil {
					continue
				}
				res += n1 * n2
			}
		}
	}

	return res
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func evaluateWithDo(s string) int {
	n := len(s)
	res := 0
	fmt.Println("Length before", n)
	s = modifyString(s)
	fmt.Println("Length after", len(s))
	n = len(s)
	for i := 0; i < n-3; i++ {
		if s[i] == 'm' && s[i+1] == 'u' && s[i+2] == 'l' && s[i+3] == '(' {
			num1 := ""
			isNum1 := true
			num2 := ""
			isNum2 := true
			for j := i + 4; j < n; j++ {
				if isDigit(s[j]) {
					num1 += string(s[j])
				} else if s[j] == ',' {
					break
				} else {
					isNum1 = false
					break
				}
			}

			if isNum1 {
				for j := i + 4 + len(num1) + 1; j < n; j++ {
					if isDigit(s[j]) {
						num2 += string(s[j])
					} else if s[j] == ')' {
						break
					} else {
						isNum2 = false
						break
					}
				}
			}

			if isNum1 && isNum2 {
				n1, err := strconv.Atoi(num1)
				if err != nil {
					continue
				}
				n2, err := strconv.Atoi(num2)
				if err != nil {
					continue
				}
				res += n1 * n2
			}
		}
	}

	return res
}

func modifyString(s string) string {
	newStr := ""
	for i := 0; i < len(s); i++ {
		if s[i] == 'd' && s[i+1] == 'o' && s[i+2] == 'n' && s[i+3] == '\'' && s[i+4] == 't' && s[i+5] == '(' && s[i+6] == ')' {
			doFound := false
			j := i + 7
			for j < len(s)-2 {
				if s[j] == 'd' && s[j+1] == 'o' && s[j+2] == '(' {
					doFound = true
					break
				}
				j++
			}
			if !doFound {
				break
			}
			i = j + 2
		} else {
			newStr += string(s[i])
		}
	}
	return newStr
}
