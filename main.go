package main

import (
	"fmt"

	"github.com/glennruffner/Demo2/workers"
)

func main() {
	// a := twoSum.TwoSum("123456733", "22")
	// fmt.Println(a)

	rawInputs := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40}
	workerCount := 3
	output, err := workers.ConvToString(rawInputs, workerCount)
	if err != nil {
		fmt.Println("Fatal Error:", err)
	}

	fmt.Println("Raw Inputs:", rawInputs)
	fmt.Println("Outputs:", output)
}
