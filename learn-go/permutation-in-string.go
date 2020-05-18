/**
Given two strings s1 and s2, write a function to return true if s2 contains the permutation of s1.
In other words, one of the first string's permutations is the substring of the second string.

Example 1:
Input: s1 = "ab" s2 = "eidbaooo"
Output: True
Explanation: s2 contains one permutation of s1 ("ba").
*/

package exercise

func checkInclusion(s1 string, s2 string) bool {
	result := false
	length := len(s2) - len(s1) + 1

	if len(s2) < len(s1) {
		return result
	}

	var s1_count, s2_count [26]int

	for i := 0; i < len(s1); i++ {
		s1_count[s1[i]-'a']++
		s2_count[s2[i]-'a']++
	}

	if s1_count == s2_count {
		return true
	}

	slider := func(left int, right int, i int) bool {
		s2_count[s2[left]-'a'] -= 1
		s2_count[s2[right]-'a'] += 1
		if s1_count == s2_count {
			return true
		}
		return false
	}

	for i := 1; i < length; i++ {
		if slider(i-1, i+len(s1)-1, i) {
			return true
		}
	}

	return result

}
