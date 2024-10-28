package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(SwapMaxMin(231))
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
			curr := JoinInt(SwapValues(digits, i, j))
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

func SwapValues[T comparable](s []T, i int, j int) []T {
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
