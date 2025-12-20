package sort_search

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	nums := []int{3, 5, 4, 2, 1, 3, 5, 6, 7, 34, 2, 1, 45, 6, 7, 43, 23, 1, 4, 7, 65}
	temp := make([]int, len(nums))
	mergeSort(nums, temp, 0, len(nums)-1)
	for _, v := range nums {
		fmt.Println(v)
	}

}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums := []int{3, 5, 4, 2, 1, 3, 5, 6, 7, 34, 2, 1, 45, 6, 7, 43, 23, 1, 4, 7, 65}
		temp := make([]int, len(nums))
		mergeSort(nums, temp, 0, len(nums)-1)
	}
}

func mergeSort(nums, temp []int, left, right int) {
	if left < right {
		mid := (left + right) / 2
		mergeSort(nums, temp, left, mid)
		mergeSort(nums, temp, mid+1, right)
		mergeSortMerge(nums, temp, left, mid, right)
	}
}

func mergeSortMerge(nums, temp []int, left, mid, right int) {
	// 两个合并
	p1 := left
	p2 := mid + 1
	k := left

	for p1 <= mid && p2 <= right {
		if nums[p1] < nums[p2] {
			temp[k] = nums[p1]
			k++
			p1++
		} else {
			temp[k] = nums[p2]
			k++
			p2++
		}
	}

	for p1 <= mid {
		temp[k] = nums[p1]
		k++
		p1++
	}

	for p2 <= right {
		temp[k] = nums[p2]
		k++
		p2++
	}

	for i := left; i <= right; i++ {
		nums[i] = temp[i]
	}

}
