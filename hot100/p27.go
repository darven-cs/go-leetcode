package hot100

//	func removeElement(nums []int, val int) int {
//		// 慢指针如果是val就停下,等待快指针找指给他切换
//		length := len(nums)
//
//		slow := length - 1
//		for fast := 0; fast < slow && fast < length; fast++ {
//			if nums[slow] == val {
//				slow--
//			} else if nums[slow] != val && nums[fast] == val {
//				nums[fast] = nums[slow]
//				slow--
//			}
//		}
//
//		return slow + 1
//	}
func removeElement(nums []int, val int) int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}
