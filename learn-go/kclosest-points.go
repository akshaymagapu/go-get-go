/**
We have a list of points on the plane.  Find the K closest points to the origin (0, 0).

(Here, the distance between two points on a plane is the Euclidean distance.)

You may return the answer in any order.  The answer is guaranteed to be unique (except for the order that it is in.)
*/
package exercise

import "sort"

func kClosest(points [][]int, K int) [][]int {
	sort.Slice(points, func(a, b int) bool {
		aVal := (points[a][0] * points[a][0]) + (points[a][1] * points[a][1])
		bVal := (points[b][0] * points[b][0]) + (points[b][1] * points[b][1])
		return aVal < bVal
	})
	return points[:K]
}
