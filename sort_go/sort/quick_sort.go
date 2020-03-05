package sort

import (
	"fmt"
)

func partition(nums []int, lo, hi int) int {
	start, end, key := lo, hi, nums[lo]
	for start < end {
		for start < hi && nums[start] <= key {
			start++
		}
		for end > lo && nums[end] > key {
			end--
		}
		if start < end {
			nums[start], nums[end] = nums[end], nums[start]
		} else {
			fmt.Println(start, end)
		}
	}
	nums[lo], nums[end] = nums[end], nums[lo]
	return end
}

func _quickSort(nums []int, lo, hi int) {
	if hi <= lo {
		return
	}
	j := partition(nums, lo, hi)
	_quickSort(nums, lo, j-1) // nums[lo,,,j-1]
	_quickSort(nums, j+1, hi) // nums[j+1,,,hi]

}

// QuickSort 递归实现快排
func QuickSort(nums []int) {
	_quickSort(nums, 0, len(nums)-1)
}
