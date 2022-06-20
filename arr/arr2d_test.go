package arr

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TraversalDfs(t *testing.T) {
	const SZ = 3
	t.Run(t.Name(), func(t *testing.T) {
		m := makeMatrix(SZ, SZ, func(i, j int) int { return SZ*i + j })
		result := traversalDfs(m)
		assert.Equal(t, []int{0, 1, 2, 5, 8, 7, 4, 3, 6}, result)
	})
}

func Test_TraversalBfs(t *testing.T) {
	const SZ = 3
	t.Run(t.Name(), func(t *testing.T) {
		m := makeMatrix(SZ, SZ, func(i, j int) int { return SZ*i + j })
		result := traversalBfs(m)
		assert.Equal(t, []int{0, 1, 3, 2, 4, 6, 5, 7, 8}, result)
	})
}

func Test_DijkstraShortestPath(t *testing.T) {
	var ts = []struct {
		input            [][]Edge // adjacent list
		start, goal, exp uint
	}{{[][]Edge{
		{{node: 2, cost: 10}, {node: 1, cost: 1}},
		{{node: 3, cost: 2}},
		{{node: 1, cost: 2}, {node: 3, cost: 3}, {node: 4, cost: 1}},
		{{node: 0, cost: 7}, {node: 4, cost: 2}},
		{},
	}, 0, 4, 5}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			if result, ok := dijkstraShortestPath(ts[i].input, ts[i].start, ts[i].goal); ok {
				assert.Equal(t, result, ts[i].exp)
			} else {
				t.Errorf("not found path from %d to %d for input %v", ts[i].start, ts[i].goal, ts[i].input)
			}
		})
	}
}

func Test_NumIslands(t *testing.T) {
	var ts = []struct {
		input [][]byte
		exp   int
	}{{[][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'}}, 3}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := numIslands(ts[i].input)
			assert.Equal(t, ts[i].exp, result)
		})
	}

}
