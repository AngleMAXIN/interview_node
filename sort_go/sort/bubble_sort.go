package sort

// Bubble 冒泡排序
func Bubble(nums []int, n int) []int {
	if n <= 1 {
		return nums
	}
	for j := n; j >= 1; j-- {
		for i := 1; i < j; i++ {
			if nums[i] < nums[i-1] {
				nums[i], nums[i-1] = nums[i-1], nums[i]
			}
		}
	}
	return nums
}
