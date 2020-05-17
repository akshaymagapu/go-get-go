/**
Given a string s and a non-empty string p, find all the start indices of p's anagrams in s.

Strings consists of lowercase English letters only and the length of both strings s and p will not be larger than 20,100.

The order of output does not matter.

Example 1:

Input:
s: "cbaebabacd" p: "abc"
Output:
[0, 6]
Explanation:
The substring with start index = 0 is "cba", which is an anagram of "abc".
The substring with start index = 6 is "bac", which is an anagram of "abc".
*/

package exercise

func findAnagrams(s string, p string) []int {
	result := []int{}
	if len(s) < len(p) {
		return result
	}
	length := len(s) - len(p) + 1

	var p_count, s_count [26]int

	for i := 0; i < len(p); i++ {
		p_count[p[i]-'a']++
		s_count[s[i]-'a']++
	}

	if p_count == s_count {
		result = append(result, 0)
	}

	slider := func(left int, right int, i int) {
		s_count[s[left]-'a'] -= 1
		s_count[s[right]-'a'] += 1
		if p_count == s_count {
			result = append(result, i)
		}
	}

	for i := 1; i < length; i++ {
		slider(i-1, i+len(p)-1, i)
	}

	return result

}
