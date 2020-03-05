package main

import (
	"fmt"
	"math/rand"
	"srot_go/sort"
	"time"
)

const (
	// N slice lenght
	N = 40
	// Limit slice values range
	Limit = 120
)

func main() {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0, N)
	for i := 0; i < N; i++ {
		nums = append(nums, rand.Intn(Limit))
	}
	fmt.Println("before:", nums)
	// sort.QuickSort(nums)
	sort.FindSort(nums)
	fmt.Println("after:", nums)

	for i := 1; i < N; i++ {
		if nums[i] < nums[i-1] {
			fmt.Println("error", nums[i], nums[i-1])
		}
	}
}
