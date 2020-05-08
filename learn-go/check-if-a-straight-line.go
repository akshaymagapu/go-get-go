/**
You are given an array coordinates, coordinates[i] = [x, y], where [x, y] represents the coordinate of a point.
Check if these points make a straight line in the XY plane.
Input: coordinates = [[1,2],[2,3],[3,4],[4,5],[5,6],[6,7]]
Output: true
*/

package exercise

func findSlope(pointA []int, pointB []int) float64 {
	if pointB[0]-pointA[0] == 0 {
		return 0
	}
	return float64(pointB[1]-pointA[1]) / float64(pointB[0]-pointA[0])
}

func checkStraightLine(coordinates [][]int) bool {
	coordinateLen := len(coordinates)

	if coordinateLen < 2 {
		return false
	}
	if coordinateLen == 2 {
		return true
	}

	lineSlope := findSlope(coordinates[0], coordinates[1])

	for i := 2; i < coordinateLen; i++ {
		currentSlope := findSlope(coordinates[i-1], coordinates[i])
		if lineSlope != currentSlope {
			return false
		}
	}

	return true
}
