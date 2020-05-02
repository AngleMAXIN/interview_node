package topk

import (
	"fmt"
)

func partions(nums []int, lo, hi int) int {
	start, end, key := lo, hi, nums[lo]

	for start < end {
		for start < hi && nums[start] > key {
			start++
		}
		for end > lo && nums[end] <= key {
			end--
		}
		if start < end {
			nums[start], nums[end] = nums[end], nums[start]
		}
	}
	nums[lo], nums[end] = nums[end], nums[lo]
	return start
}

func quickSort(nums []int, l, r, k int) {
	if l >= r {
		return
	}
	j := partions(nums, l, r)
	if j == k {
		return
	}
	quickSort(nums, l, j-1, k)
	quickSort(nums, j+1, r, k)
}

// TopK return top k nums
func TopK(nums []int, k int) []int {
	l := len(nums)
	fmt.Println("start: ", nums)

	quickSort(nums, 0, l-1, k)
	fmt.Println("end: ", nums)

	return nums[:k]
}
