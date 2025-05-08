package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readStdInNums() []int {
	scanner := bufio.NewScanner(os.Stdin)
	nums := []int{}
	for scanner.Scan() {
		strNum := scanner.Text()
		if len(strNum) == 0 {
			return nums
		}

		num, err := strconv.Atoi(strNum)
		if err != nil {
			panic("String to Integer conversion failed")
		}
		nums = append(nums, num)
	}
	return nums
}

func isEven(num int) bool {
	return (num & 1) == 0
}

func IsPrime(num int) bool {
	// definitely isn't prime if the number is even
	if isEven(num) {
		return false
	}

	return true
}

func main() {
	nums := readStdInNums()
	for _, num := range nums {
		fmt.Println(IsPrime(num))
	}
}
