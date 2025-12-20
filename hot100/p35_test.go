package hot100

func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums)
	for l < r {
		mid := l + (r-l)>>1
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}
