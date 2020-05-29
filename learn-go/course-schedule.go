/**
There are a total of numCourses courses you have to take, labeled from 0 to numCourses-1.

Some courses may have prerequisites, for example to take course 0 you have to first take course 1, which is expressed as a pair: [0,1]

Given the total number of courses and a list of prerequisite pairs, is it possible for you to finish all courses?

Example 1:
Input: numCourses = 2, prerequisites = [[1,0]]
Output: true
Explanation: There are a total of 2 courses to take.
To take course 1 you should have finished course 0. So it is possible.
*/
package exercise

func hasCycle(course int, cycle map[int]int, graph map[int][]int) bool {
	if cycle[course] == 2 {
		return false
	}
	if cycle[course] == 1 {
		return true
	}
	cycle[course] = 1
	if len(graph[course]) > 0 {
		for _, val := range graph[course] {
			if hasCycle(val, cycle, graph) {
				return false
			}
		}
	}
	cycle[course] = 2
	return false
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int)
	cycle := make(map[int]int)

	for _, courses := range prerequisites {
		actual := courses[0]
		dependent := courses[1]
		graph[actual] = append(graph[actual], dependent)
	}
	for key, _ := range graph {
		if hasCycle(key, cycle, graph) {
			return false
		}
	}

	return true
}
