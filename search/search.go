package search

import "fmt"

// https://leetcode.com/problems/kth-largest-element-in-an-array/
//

func partition(v []int, l, r int) int {
	fmt.Printf("0) %v/%v  %v\n", l, r, v)
	pivot := v[r]
	p_index := l
	for j := l; j < r; j++ {
		if v[j] <= pivot {
			v[j], v[p_index] = v[p_index], v[j]
			fmt.Printf("1) %v/%v/%v %v\n", v[j], v[p_index], pivot, v)
			p_index++
		}
	}
	v[p_index], v[r] = v[r], v[p_index]
	fmt.Printf("2) %v/%v %v\n", p_index, r, v)
	return p_index
}

func qsort(nums []int, l, r int) {
	if l < r {
		p := partition(nums, l, r)
		qsort(nums, l, p-1)
		qsort(nums, p+1, r)
	}
}

func FindKthLargest(nums []int, k int) int {
	qsort(nums, 0, len(nums)-1)
	return nums[len(nums)-k]
}

// https://leetcode.com/problems/kth-largest-element-in-an-array/
// quick select approach
//

func quick_select(v []int, l, r, k int) int {
	fpartition := func(v []int, l, r int) int {
		i := l
		for j := l; j <= r; j++ {
			if v[j] <= v[r] {
				v[i], v[j] = v[j], v[i]
				i += 1
			}
		}
		return i - 1
	}

	p_idx := fpartition(v, l, r)
	if k < p_idx {
		return quick_select(v, l, p_idx-1, k)
	}
	if k > p_idx {
		return quick_select(v, p_idx+1, r, k)
	}
	return v[k]
}

func FindKthLargestQselect(nums []int, k int) int {
	return quick_select(nums, 0, len(nums)-1, len(nums)-k)
}

// ---  Search range
// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
//
func SearchRange(nums []int, target int) []int {
	fbin_search := func(v []int, l, r, target int) int {
		for mid := (r + l) / 2; l <= r; mid = (r + l) / 2 {
			mid_val := v[mid]
			if mid_val < target {
				l = mid + 1
			} else if mid_val > target {
				r = mid - 1
			} else {
				return mid
			}
		}
		return -1
	}

	switch len(nums) {
	case 1:
		if nums[0] == target {
			return []int{0, 0}
		}
		fallthrough
	case 0:
		return []int{-1, -1}
	}

	first_pos := fbin_search(nums, 0, len(nums)-1, target)
	if first_pos == -1 {
		return []int{-1, -1}
	}

	start_pos, t1 := first_pos, -1
	for start_pos != -1 {
		if t1 = start_pos; start_pos > 0 {
			start_pos = fbin_search(nums, 0, start_pos-1, target)
		} else {
			start_pos = -1
		}
	}

	end_pos, t2 := t1, -1
	for end_pos != -1 {
		t2 = end_pos
		end_pos = fbin_search(nums, end_pos+1, len(nums)-1, target)
	}
	return []int{t1, t2}
}
