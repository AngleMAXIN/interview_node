package search

// func partition(nums []int, lo, hi int) int {
// 	start, end, key := lo, hi, nums[lo]
// 	fmt.Println("---", nums)
// 	for start < end {
// 		for start < hi && nums[start] <= key {
// 			start++
// 		}
// 		for end > lo && nums[end] > key {
// 			end--
// 		}
// 		if start < end {
// 			nums[start], nums[end] = nums[end], nums[start]
// 		}
// 	}
// 	nums[lo], nums[end] = nums[end], nums[lo]
// 	return end
// }

func getPartition(arr []int, start, end int) int {

	tmp := arr[start]
	for start < end {
		for start < end && arr[end] >= tmp {
			end--
		}
		arr[start] = arr[end]
		for start < end && arr[start] <= tmp {
			start++
		}
		arr[end] = arr[start]
	}
	arr[start] = tmp
	return start
}

func getLeastNumbers(nums []int, k int) []int {
	if k > len(nums) || k <= 0 {
		return nil
	}

	start, end := 0, len(nums)-1
	index := getPartition(nums, start, end)
	for index != k-1 {
		if index > k-1 {
			index = getPartition(nums, start, index-1)
		} else {
			index = getPartition(nums, index+1, end)
		}
	}
	return nums[:index+1]
}