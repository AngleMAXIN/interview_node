package search

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
		}
	}
	nums[lo], nums[end] = nums[end], nums[lo]
	return end
}

func _quickSort(nums []int, lo, hi, k int) {
	if hi <= lo {
		return
	}
	j := partition(nums, lo, hi)
	if j == k-1 {
		return
	} else if j > k-1 {
		_quickSort(nums, lo, j-1, k) // nums[lo,,,j-1]
	} else {
		_quickSort(nums, j+1, hi, k) // nums[j+1,,,hi]
	}
}

// QuickSort 递归实现快排
func QuickSort(nums []int, k int) []int {
	_quickSort(nums, 0, len(nums)-1, k)
	return nums[:k]
}
