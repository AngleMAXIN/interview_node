package main

import (
	"math"
)

func printNumbers1(n int) []int {
	max := int(math.Pow(10, float64(n)))
	res := make([]int, max)
	for i := 0; i < max; i++ {
		res[i] = i
	}
	return res
}

func majorityElement(nums []int) int {
	m, count := 0, 0
	for _, v := range nums {
		if count == 0 {
			m = v
		}
		if m == v {
			count++
		} else {
			count--
		}
	}
	return m
}
