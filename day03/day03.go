package day03

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	DEBUG1 bool = false
	DEBUG2 bool = false
)

func Part1(input string) int {
	lines := parse(input)
	sum := 0

	for _, line := range lines {
		digits := splitDigits(line)

		first, idx := maxIdx(digits[:len(digits)-1])
		second, _ := maxIdx(digits[idx+1:])
		sum += first*10 + second
		if DEBUG1 {
			fmt.Printf("Num: %v%v\n", first, second)
			fmt.Println("Sum: ", sum)
		}
	}

	return sum
}

func Part2(input string) int {
	lines := parse(input)
	sum := 0

	for _, line := range lines {
		digits := splitDigits(line)
		var bigNums []int

		for j := 12; j > 0; j-- {
			num, idx := maxIdx(digits[:len(digits)-j+1])
			bigNums = append(bigNums, num)
			digits = digits[idx+1:]
		}

		sBigNums := make([]string, len(bigNums))
		for i, v := range bigNums {
			sBigNums[i] = strconv.Itoa(v)
		}
		s := strings.Join(sBigNums, "")

		bigNum, _ := strconv.Atoi(s)
		sum += bigNum

		if DEBUG2 {
			fmt.Println("Battery: ", line)
			fmt.Println("Num: ", bigNum)
			fmt.Println("Sum: ", sum)
		}
	}

	return sum
}

func parse(input string) []string {
	input = strings.TrimSpace(input)
	if input == "" {
		return nil
	}

	return strings.Split(input, "\n")
}

func splitDigits(s string) []int {
	if s == "0" {
		return []int{0}
	}

	var digits []int
	for _, char := range s {
		digit, _ := strconv.Atoi(string(char))
		digits = append(digits, digit)
	}
	return digits
}

func maxIdx(nums []int) (max int, idx int) {
	if len(nums) == 0 {
		return 0, -1
	}

	max = nums[0]
	idx = 0

	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			idx = i
		}
	}

	return max, idx
}
