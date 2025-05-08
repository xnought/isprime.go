package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readStdinNums() []int {
	scanner := bufio.NewScanner(os.Stdin)
	nums := []int{}
	for scanner.Scan() {
		strNum := scanner.Text()
		num, err := strconv.Atoi(strNum)
		if err != nil {
			panic("String to Integer conversion failed")
		}
		nums = append(nums, num)
	}
	return nums
}

func main() {
	nums := readStdinNums()
	fmt.Println(nums)
}
