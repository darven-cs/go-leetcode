package hot100

import "testing"

func TestName(t *testing.T) {
	findOriginalArray([]int{6, 3, 0, 1})
}

func findOriginalArray(changed []int) []int {
	if len(changed)%2 != 0 {
		return []int{}
	}

	// slices.Sort(changed)
	m := map[int]bool{}
	for _, v := range changed {
		m[v] = false
	}
	num := 0
	ans := []int{}
	for k, _ := range m {
		if k == 0 {
			continue
		}
		t := k * 2

		if v, ok := m[t]; ok && !v {
			num += 1
			ans = append(ans, k)
			m[t] = true
			m[k] = true
		}
	}
	if num == (len(changed) / 2) {
		return ans
	} else {
		return []int{}
	}
}
