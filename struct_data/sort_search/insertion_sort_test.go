package sort_search

import (
	"fmt"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	nums := []int{3, 5, 4, 2, 1, 3, 5, 6, 7, 34, 2, 1, 45, 6, 7, 43, 23, 1, 4, 7, 65}
	length := len(nums)
	for i := 1; i < length; i++ {
		temp := nums[i]
		j := i - 1
		for ; j >= 0; j-- {
			if nums[j] > temp {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}
		nums[j+1] = temp
	}

	for _, v := range nums {
		fmt.Printf("%d ", v)
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums := []int{3, 5, 4, 2, 1, 3, 5, 6, 7, 34, 2, 1, 45, 6, 7, 43, 23, 1, 4, 7, 65}
		length := len(nums)
		for i := 1; i < length; i++ {
			temp := nums[i]
			j := i - 1
			for ; j >= 0; j-- {
				if nums[j] > temp {
					nums[j+1] = nums[j]
				} else {
					break
				}
			}
			nums[j+1] = temp
		}
	}
}
