package sort_search

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	nums := []int{3, 5, 4, 2, 1, 3, 5, 6, 7, 34, 2, 1, 45, 6, 7, 43, 23, 1, 4, 7, 65}
	length := len(nums)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i] > nums[j] {
				swapElements(nums, i, j)
			}
		}
	}

	for _, v := range nums {
		fmt.Printf("%d ", v)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums := []int{3, 5, 4, 2, 1, 3, 5, 6, 7, 34, 2, 1, 45, 6, 7, 43, 23, 1, 4, 7, 65}
		length := len(nums)
		for i := 0; i < length-1; i++ {
			for j := i + 1; j < length; j++ {
				if nums[i] > nums[j] {
					swapElements(nums, i, j)
				}
			}
		}
	}
}

func swapElements(nums []int, i, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}
