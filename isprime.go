package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readStdInNums() []uint64 {
	scanner := bufio.NewScanner(os.Stdin)
	nums := []uint64{}
	for scanner.Scan() {
		strNum := scanner.Text()
		if len(strNum) == 0 {
			return nums
		}

		num, err := strconv.ParseUint(strNum, 10, 64)
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}
	return nums
}

func isEven(num uint64) bool {
	return (num & 1) == 0
}

func IsPrime(num uint64) bool {
	// we consider 1 prime
	if num == 1 {
		return true
	}
	// definitely isn't prime if the number is even (but 2 is only prime even)
	if num != 2 && isEven(num) {
		return false
	}

	// check if any numbers divide into num
	// if yes, not prime!
	for i := uint64(2); i < num; i++ {
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
