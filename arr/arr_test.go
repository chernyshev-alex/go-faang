package arr

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_twoSum(t *testing.T) {
	var ts = []struct {
		input    []int
		target   int
		expected [2]int
	}{{[]int{2, 7, 11, 15}, 9, [2]int{0, 1}}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			var result [2]int
			copy(result[:], twoSum(ts[i].input, ts[i].target))
			if result != ts[i].expected {
				t.Errorf("input : %v; exp.: %v, got: %v", ts[i].input, ts[i].expected, result)
			}
		})
	}
}

func Test_maxArea(t *testing.T) {
	var ts = []struct {
		input    []int
		expected int
	}{{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := maxArea(ts[i].input)
			if result != ts[i].expected {
				t.Errorf("input : %v; exp.: %v, got: %v", ts[i].input, ts[i].expected, result)
			}
		})
	}
}

func Test_trap(t *testing.T) {
	var ts = []struct {
		input    []int
		expected int
	}{{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := trap(ts[i].input)
			if result != ts[i].expected {
				t.Errorf("input : %v; exp.: %v, got: %v", ts[i].input, ts[i].expected, result)
			}
		})
	}
}

func Test_fibo(t *testing.T) {
	t.Run("fibo", func(t *testing.T) {
		assert.Equal(t, []int64{1, 2, 3, 5, 8, 13, 21}, fibo_gen(6))
	})
}

func Test_binarySearch(t *testing.T) {
	t.Run("fibo", func(t *testing.T) {
		arr := []int{12, 81, 41, 33, 99, 55, 78, 25, 9, 35}
		assert.Equal(t, 8, binarySearch(arr, 81))
		fmt.Println("src arrr", arr)
	})
}

func Bench_binarySearch(t *testing.T) {
	t.Run("fibo", func(t *testing.T) {
		arr := []int{12, 81, 41, 33, 99, 55, 78, 25, 9, 35}
		assert.Equal(t, 8, binarySearch(arr, 81))
		fmt.Println("src arrr", arr)
	})
}

func BenchmarkBinarySearch(b *testing.B) {
	b.ReportAllocs()
	arr := []int{12, 81, 41, 33, 99, 55, 78, 25, 9, 35}
	for i := 0; i < b.N; i++ {
		binarySearch(arr, 81)
	}
}

// BenchmarkBinnarySearch-4   	16125369	        74.10 ns/op	      24 B/op	       1 allocs/op

func BenchmarkBinarySearch2(b *testing.B) {
	b.ReportAllocs()
	arr := []int{12, 81, 41, 33, 99, 55, 78, 25, 9, 35}
	for i := 0; i < b.N; i++ {
		binarySearch2(arr, 81)
	}
}

//cpu: Intel(R) Core(TM) i5-7360U CPU @ 2.30GHz
// BenchmarkBinnarySearch2-4   	16233514	        73.97 ns/op	      24 B/op	       1 allocs/op
