package hot100

import "testing"

// 单调递增
func dailyTemperatures(temperatures []int) []int {
	length := len(temperatures)
	answer := make([]int, length)
	stack := make([]int, 0)

	for i := 0; i < length; i++ {
		// 如果大的话要把小的剔除,然后记录坐标
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			preIndex := stack[len(stack)-1]

			stack = stack[:len(stack)-1]
			answer[preIndex] = i - preIndex
		}
		stack = append(stack, i)
	}
	return answer
}

func TestDailyTemperatures(t *testing.T) {
	// 测试用例1: 基本情况
	temperatures1 := []int{73, 74, 75, 71, 69, 72, 76, 73}
	expected1 := []int{1, 1, 4, 2, 1, 1, 0, 0}
	result1 := dailyTemperatures(temperatures1)
	for i, v := range result1 {
		if v != expected1[i] {
			t.Errorf("Test case 1 failed: expected %v, got %v", expected1, result1)
			break
		}
	}

	// 测试用例2: 递减序列
	temperatures2 := []int{30, 20, 10}
	expected2 := []int{0, 0, 0}
	result2 := dailyTemperatures(temperatures2)
	for i, v := range result2 {
		if v != expected2[i] {
			t.Errorf("Test case 2 failed: expected %v, got %v", expected2, result2)
			break
		}
	}

	// 测试用例3: 递增序列
	temperatures3 := []int{10, 20, 30}
	expected3 := []int{1, 1, 0}
	result3 := dailyTemperatures(temperatures3)
	for i, v := range result3 {
		if v != expected3[i] {
			t.Errorf("Test case 3 failed: expected %v, got %v", expected3, result3)
			break
		}
	}

	// 测试用例4: 相同温度
	temperatures4 := []int{30, 30, 30}
	expected4 := []int{0, 0, 0}
	result4 := dailyTemperatures(temperatures4)
	for i, v := range result4 {
		if v != expected4[i] {
			t.Errorf("Test case 4 failed: expected %v, got %v", expected4, result4)
			break
		}
	}

	// 测试用例5: 单个元素
	temperatures5 := []int{30}
	expected5 := []int{0}
	result5 := dailyTemperatures(temperatures5)
	for i, v := range result5 {
		if v != expected5[i] {
			t.Errorf("Test case 5 failed: expected %v, got %v", expected5, result5)
			break
		}
	}
}
