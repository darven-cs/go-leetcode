package hot100

import "sort"

// a+b+c=0 也就是 -a=b+c
// 相当于如果没有负数就提前结束就行
// i指针是固定
// j指针和k指针找组合
// 为了确保不重复就遍历前后,去重
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	length := len(nums)

	ans := make([][]int, 0)

	for i := 0; i < length; i++ {
		if nums[i] > 0 {
			break
		}
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}

		target := -nums[i]

		j := i + 1
		k := length - 1

		for j < k {
			if nums[j]+nums[k] == target {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				// 跳过重复元素
				for j < k && nums[j] == nums[j+1] {
					j++
				}
				for j < k && nums[k] == nums[k-1] {
					k--
				}
				j++
				k--
			} else if nums[j]+nums[k] > target {
				k--
			} else {
				j++
			}
		}
	}
	return ans
}
