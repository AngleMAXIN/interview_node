package windows

func maxSlidingWindow(nums []int, k int) []int {
	l := len(nums)
	if l < 1 {
		return nil
	}
	windows, res := []int{}, []int{}
	winLen := 0
	for i, v := range nums {
		// 判断窗口大小是否越界，是，则弹出最左侧数据
		if i >= k && windows[0] <= i-k {
			windows = windows[1:]
		}
		
		// 维护队列中的大小，保持最左侧的永远是大的
		winLen = len(windows)
		for winLen > 0 && nums[windows[winLen-1]] <= v {
			windows = windows[:winLen-1]
			winLen--

		}
		windows = append(windows, i)
		// 如果窗口已经满了，开始往res里存
		if i >= k-1 {
			res = append(res, nums[windows[0]])
		}
	}
	return res
}
