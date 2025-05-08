package main

import (
	"bufio"
	"fmt"
	"math"
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

func isPrime(num uint64) bool {
	// 0 and 1 are not prime
	if num == 0 || num == 1 {
		return false
	}

	// 2 and 3 are prime
	if num <= 3 {
		return true
	}

	// definitely isn't prime if the number is even
	if isEven(num) {
		return false
	}

	//
	// NOW EXHAUSTIVELY SEARCH
	//

	// check if any numbers divide into num
	// if yes, not prime!
	// only need to check up to sqrt of num
	// and since already checked even, don't even have to try even divisors
	for i := uint64(3); i < uint64(math.Sqrt(float64(num))); i += 2 {
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
		if isPrime(num) {
			fmt.Println("1")
		} else {
			fmt.Println("0")
		}
	}
}
