package sort

var aux []int

func less(a, b int) bool {
	if a >= b {
		return false
	}
	return true
}

func merge(nums []int, lo, mid, hi int) {
	i, j := lo, mid+1
	for k := 0; k <= hi; k++ {
		aux[k] = nums[k]
	}

	for k := lo; k <= hi; k++ {
		// fmt.Println(len(nums), cap(nums))
		if i > mid {
			nums[k] = aux[j]
			j++
		} else if j > hi {
			nums[k] = aux[i]
			i++
		} else if less(aux[j], aux[i]) {
			nums[k] = aux[j]
			j++
		} else {
			nums[k] = aux[i]
			i++
		}
	}
	// fmt.Println(nums[lo:hi])

}

// MergeSort 归并排序
func MergeSort(nums []int) {
	aux = make([]int, len(nums))
	_mergeSort(nums, 0, len(nums)-1)
}

func _mergeSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	mid := l + (r-l)/2
	_mergeSort(nums, l, mid)
	_mergeSort(nums, mid+1, r)
	merge(nums, l, mid, r)
}
