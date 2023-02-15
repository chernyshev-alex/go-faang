package arr

// utils

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Tasks

// given array of ints, return indices of the two numbers that add up to a given target
// https://leetcode.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i1, v := range nums {
		if i2, ok := m[target-v]; ok {
			return []int{i2, i1}
		}
		m[v] = i1
	}
	return make([]int, 2)
}

// https://leetcode.com/problems/container-with-most-water/

func maxArea(height []int) int {
	var best_area, x1, x2 int = 0, 0, len(height) - 1
	for x1 < x2 {
		area := min(height[x1], height[x2]) * (x2 - x1)
		best_area = max(area, best_area)
		if height[x1] <= height[x2] {
			x1 += 1
		} else {
			x2 -= 1
		}
	}
	return best_area
}

// https://leetcode.com/problems/trapping-rain-water/

func trap(height []int) int {
	result := 0
	var li, ri, lmax, rmax = 0, len(height) - 1, 0, 0
	for li < ri {
		if height[li] <= height[ri] {
			if height[li] >= lmax {
				lmax = height[li]
			} else {
				result += lmax - height[li]
			}
			li += 1
		} else {
			if height[ri] >= rmax {
				rmax = height[ri]
			} else {
				result += rmax - height[ri]
			}
			ri -= 1
		}
	}
	return result
}

func fibo_gen(n int) []int64 {
	var result = make([]int64, 0)

	var fiboFun func(int, int64, int64)

	fiboFun = func(n int, a, b int64) {
		result = append(result, a+b)
		if n > 0 {
			fiboFun(n-1, b, a+b)
		}
	}
	fiboFun(n, 0, 1)
	return result

}
