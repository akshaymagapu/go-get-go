/**
Given a set of N people (numbered 1, 2, ..., N), we would like to split everyone into two groups of any size.

Each person may dislike some other people, and they should not go into the same group.

Formally, if dislikes[i] = [a, b], it means it is not allowed to put the people numbered a and b into the same group.

Return true if and only if it is possible to split everyone into two groups in this way.

Example 1:

Input: N = 4, dislikes = [[1,2],[1,3],[2,4]]
Output: true
Explanation: group1 [1,4], group2 [2,3]
*/
package exercise

func possibleBipartition(N int, dislikes [][]int) bool {
	graph := make([][]bool, N+1)
	color := make([]int, N+1)
	var dfs func(value int, i int) bool

	for i := 1; i <= N; i++ {
		graph[i] = make([]bool, N+1)
	}

	for _, d := range dislikes {
		firstEle, secondEle := d[0], d[1]
		graph[firstEle][secondEle] = true
		graph[secondEle][firstEle] = true
	}

	dfs = func(value int, i int) bool {
		if color[i] != 0 {
			return color[i] == value
		}
		color[i] = value
		for j := 1; j <= N; j++ {
			if graph[i][j] {
				if !dfs(-value, j) {
					return false
				}
			}
		}
		return true
	}

	for i := 1; i <= N; i++ {
		if color[i] == 0 {
			if !dfs(1, i) {
				return false
			}
		}
	}
	return true
}
