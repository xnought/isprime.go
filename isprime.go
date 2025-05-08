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
	// definitely isn't prime if the number is even (but 2 is only prime even)
	if num != 2 && isEven(num) {
		return false
	}

	// check if any numbers divide into num
	// if yes, not prime!
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}

	// if no, prime!
	return true
}

func main() {
	nums := readStdInNums()
	for _, num := range nums {
		fmt.Println(IsPrime(num))
	}
}
