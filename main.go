package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		n := 231
		fmt.Println("Using default value", n)
		fmt.Println(SwapMaxMin(n))
	} else {
		for _, arg := range os.Args[1:] {
			fmt.Println()
			fmt.Println("Finding max and min of", arg)
			n, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Could not parse number")
				continue
			}
			fmt.Println(SwapMaxMin(n))
		}
	}
}

func SwapMaxMin(n int) (max int, min int) {
	digits := SplitInt(n)
	max = n
	min = n
	for i := range digits {
		for j := i + 1; j < len(digits); j++ {
			// No leading zeroes allowed
			if i == 0 && digits[j] == 0 {
				continue
			}
			curr := JoinInt(SwapElements(digits, i, j))
			if curr > max {
				max = curr
			}
			if curr < min {
				min = curr
			}
		}
	}
	return max, min
}

func JoinInt(s []int) int {
	x := 0
	for _, y := range s {
		x = x*10 + y
	}
	return x
}

func SwapElements[T comparable](s []T, i int, j int) []T {
	output := slices.Clone(s)
	output[i] = s[j]
	output[j] = s[i]
	return output
}

func SplitInt(n int) []int {
	output := []int{}
	for {
		output = append(output, n%10)
		n = n / 10
		if n == 0 {
			break
		}
	}
	slices.Reverse(output)
	return output
}
