package hot100

import "testing"

func TestP416(t *testing.T) {
	nums := []int{3, 3, 6, 8, 16, 16, 16, 18, 20}
	partition := canPartition(nums)
	println(partition)
}

func canPartition(nums []int) bool {
	num := 0
	for _, v := range nums {
		num += v
	}

	if num%2 != 0 {
		return false
	}

	num = num / 2

	dp := make([]int, len(nums))
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		if dp[i-1]+nums[i] <= num {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = dp[i-1]
		}
	}

	return dp[len(nums)-1] == num
}
