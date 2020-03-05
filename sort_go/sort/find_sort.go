package sort

import "fmt"

func findMaxPos(nums []int, n int) int {
	max := nums[0]
	pos := 0
	for i := 0; i < n; i++ {
		if nums[i] > max {
			pos, max = i, nums[i]
		}
	}

	return pos
}

// FindSort 选择排序
func FindSort(nums []int) []int {
	n, pos := len(nums), 0
	for n >= 1 {
		pos = findMaxPos(nums, n)
		nums[pos], nums[n-1] = nums[n-1], nums[pos]
		n--
	}
	fmt.Printf("%p\n", nums)
	nums = append(nums, 5, 6, 7, 78, 7, 9, 9, 9)
	fmt.Printf("%p\n", nums)

	return nums
}
