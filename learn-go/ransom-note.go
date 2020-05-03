/**
Given an arbitrary ransom note string and another string containing letters from all the magazines,
write a function that will return true if the ransom note can be constructed from the magazines ; otherwise, it will return false.

Each letter in the magazine string can only be used once in your ransom note.

canConstruct("a", "b") -> false
canConstruct("aa", "ab") -> false
canConstruct("aa", "aab") -> true
*/

package exercise

func canConstruct(ransomNote string, magazine string) bool {

	magazineCount := make(map[rune]int)

	for _, c := range magazine {
		magazineCount[c]++
	}

	for _, c := range ransomNote {
		if magazineCount[c] > 0 {
			magazineCount[c]--
		} else {
			return false
		}
	}

	return true

}
