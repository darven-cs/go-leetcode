package hot100

import (
	"fmt"
	"testing"
)

// map实现
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	n := len(nums)

	for i := 0; i < n; i++ {
		tmp := target - nums[i]
		if index, ok := m[tmp]; ok {
			return []int{i, index}
		}
		m[nums[i]] = i
	}
	return nil
}

// 双指针
// func twoSum(nums []int, target int) []int {
// 	length := len(nums)
// 	j := length - 1
// 	for i := 0; i < len(nums); {
// 		if nums[i]+nums[j] > target {
// 			j--
// 		} else if nums[i]+nums[j] < target {
// 			i++
// 			j = length - 1
// 		} else {
// 			return []int{i, j}
// 		}
// 	}

// 	return nil
// }

func TestTwoSum(t *testing.T) {
	fmt.Println(twoSum([]int{1, 2, 3, 4}, 6))
}
