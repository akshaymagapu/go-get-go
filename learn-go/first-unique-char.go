/**
Given a string, find the first non-repeating character in it and return it's index.
If it doesn't exist, return -1.

s = "aks"
return 0.
*/

package exercise

func firstUniqChar(s string) int {

	charCount := make(map[rune]int, len(s))

	for _, c := range s {
		charCount[c]++
	}

	for i, c := range s {
		if charCount[c] == 1 {
			return i
		}
	}
	return -1
}
