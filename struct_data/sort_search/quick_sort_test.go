package sort_search

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	nums := []int{3, 5, 4, 2, 1, 3, 5, 6, 7, 34, 2, 1, 45, 6, 7, 43, 23, 1, 4, 7, 65}
	quickSort(nums, 0, len(nums)-1)

	for _, v := range nums {
		fmt.Println(v)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums := []int{3, 5, 4, 2, 1, 3, 5, 6, 7, 34, 2, 1, 45, 6, 7, 43, 23, 1, 4, 7, 65}
		quickSort(nums, 0, len(nums)-1)
	}
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}

	i := left
	j := right
	mid := nums[left]

	for i < j {
		// 先找右边
		for mid <= nums[j] && i < j {
			j--
		}
		// 再找左边
		for i < j && mid >= nums[i] {
			i++
		}
		temp := nums[i]
		nums[i] = nums[j]
		nums[j] = temp
	}
	// 将中间点放到中间
	nums[left] = nums[i]
	nums[i] = mid

	// 继续递归
	quickSort(nums, left, i-1)
	quickSort(nums, i+1, right)
}
