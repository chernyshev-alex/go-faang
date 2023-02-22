package graphs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// --- Adjacency Lists ---

func Test_TraversalBfs(t *testing.T) {
	var ts = []struct {
		input [][]int
		exp   []int
	}{{[][]int{{1, 3}, {0}, {3, 8}, {0, 2, 4, 5}, {3, 6}, {3}, {4, 7}, {6}, {2}}, []int{0, 1, 3, 2, 4, 5, 8, 6, 7}}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := traversal_bfs(ts[i].input)
			assert.Equal(t, ts[i].exp, result)
		})
	}
}

func Test_TraversalDfs(t *testing.T) {
	var ts = []struct {
		input [][]int
		exp   []int
	}{{[][]int{{1, 3}, {0}, {3, 8}, {0, 2, 4, 5}, {3, 6}, {3}, {4, 7}, {6}, {2}}, []int{0, 1, 3, 2, 8, 4, 6, 7, 5}}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			values := make([]int, 0)
			seen := make(map[int]bool)
			traversal_dfs(0, ts[i].input, &values, &seen)
			assert.Equal(t, ts[i].exp, values)
		})
	}
}

// --- Adjacency Matrix ---

func Test_TraversalBfsMatrix(t *testing.T) {
	var ts = []struct {
		input [][]int
		exp   []int
	}{{[][]int{
		{0, 1, 0, 1, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 1},
		{1, 0, 1, 0, 1, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 1, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 0}}, []int{0, 1, 3, 2, 4, 5, 8, 6, 7}}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := traversalBfsMatrix(ts[i].input)
			assert.Equal(t, ts[i].exp, result)
		})
	}
}

func Test_TraversalDfsMatrix(t *testing.T) {
	var ts = []struct {
		input [][]int
		exp   []int
	}{{[][]int{
		{0, 1, 0, 1, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 1},
		{1, 0, 1, 0, 1, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 1, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 0}}, []int{0, 1, 3, 2, 8, 4, 6, 7, 5}}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			values := make([]int, 0)
			seen := make(map[int]bool)
			traversalDfsMatrix(0, ts[i].input, &values, &seen)
			assert.Equal(t, ts[i].exp, values)
		})
	}
}

// Time Needed to Inform All Employees
// https://leetcode.com/problems/time-needed-to-inform-all-employees/
func Test_NumOfMinutes(t *testing.T) {
	var ts = []struct {
		n, headID           int
		manager, informTime []int
		exp                 int
	}{{1, 0, []int{-1}, []int{0}, 0},
		{6, 2, []int{2, 2, -1, 2, 2, 2}, []int{0, 0, 1, 0, 0, 0}, 1}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := numOfMinutes(ts[i].n, ts[i].headID, ts[i].manager, ts[i].informTime)
			assert.Equal(t, ts[i].exp, result)
		})
	}
}

// Course Schedule
// https://leetcode.com/problems/course-schedule/
//
// There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1.
// You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates
// that you must take course bi first if you want to take course ai.
//
// For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
// Return true if you can finish all courses. Otherwise, return false.
func Test_CourseScheduleCanFinish(t *testing.T) {
	var ts = []struct {
		input      [][]int
		numCourses int
		exp        bool
	}{{[][]int{{1, 0}}, 2, true}, {[][]int{{1, 0}, {0, 2}}, 2, false}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := canFinish(ts[i].numCourses, ts[i].input)
			assert.Equal(t, ts[i].exp, result)
		})
	}
}

// Network Delay Time
// https://leetcode.com/problems/network-delay-time/
// You are given a network of n nodes, labeled from 1 to n.
// You are also given times, a list of travel times as directed edges times[i] = (ui, vi, wi), where ui is the source node,
// vi is the target node, and wi is the time it takes for a signal to travel from source to target.
// We will send a signal from a given node k. Return the minimum time it takes for all the n nodes to receive the signal.
// If it is impossible for all the n nodes to receive the signal, return -1.
func Test_NetworkDelayTimeDejkstra(t *testing.T) {
	var ts = []struct {
		input [][]int
		n, k  int
		exp   int
	}{{[][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2, 2},
		{[][]int{{1, 2, 1}}, 2, 1, 1},
		{[][]int{{1, 2, 1}}, 2, 2, -1}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := networkDelayTime_Dejkstra(ts[i].input, ts[i].n, ts[i].k)
			assert.Equal(t, ts[i].exp, result)
		})
	}
}

// finding the shortest distance to all vertices from the source using the Bellman-Ford algorithm
func Test_ShortestPathBellmanFordNegativeWeigths(t *testing.T) {
	var ts = []struct {
		v, e, src int
		edges     []Edge
		dist      []int
	}{{5, 8, 0, []Edge{{0, 1, -1}, {0, 2, 4}, {1, 2, 3}, {1, 3, 2}, {1, 4, 2}, {3, 2, 5}, {3, 1, 1}, {4, 3, -3}},
		[]int{0, -1, 2, -2, 1}}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			dist, _ := shortestPathBellmanFordNegativeWeigths(ts[i].v, ts[i].e, ts[i].src, ts[i].edges)
			assert.Equal(t, ts[i].dist, dist)
		})
	}
}
func Test_NetworkDelayTimeBellmanFord(t *testing.T) {
	var ts = []struct {
		input [][]int
		n, k  int
		exp   int
	}{{[][]int{{1, 4, 2}, {1, 2, 9}, {4, 2, -4}, {2, 5, -3}, {4, 5, 6}, {3, 2, 3}, {5, 3, 7}, {3, 1, 5}}, 5, 1, 2}}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			result := networkDelayTime_BellmanFord(ts[i].input, ts[i].n, ts[i].k)
			assert.Equal(t, ts[i].exp, result)
		})
	}
}

func Test_DijkstraPrismShortestPath(t *testing.T) {

	var ts = []struct {
		src   int
		graph [][]int
		exp   []vdist
	}{
		{0, [][]int{
			{0, 4, 0, 0, 0, 0, 0, 8, 0},
			{4, 0, 8, 0, 0, 0, 0, 11, 0},
			{0, 8, 0, 7, 0, 4, 0, 0, 2},
			{0, 0, 7, 0, 9, 14, 0, 0, 0},
			{0, 0, 0, 9, 0, 10, 0, 0, 0},
			{0, 0, 4, 14, 10, 0, 2, 0, 0},
			{0, 0, 0, 0, 0, 2, 0, 1, 6},
			{8, 11, 0, 0, 0, 0, 1, 0, 7},
			{0, 0, 2, 0, 0, 0, 6, 7, 0},
		},
			[]vdist{
				{v: 0, dist: 0},
				{v: 1, dist: 4},
				{v: 2, dist: 12},
				{v: 3, dist: 19},
				{v: 4, dist: 21},
				{v: 5, dist: 11},
				{v: 6, dist: 9},
				{v: 7, dist: 8},
				{v: 8, dist: 14},
			},
		},
	}

	for i := range ts {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			assert.Equal(t, ts[i].exp, findShortestPathDijkstraPrism(ts[i].src, ts[i].graph))
		})
	}
}
